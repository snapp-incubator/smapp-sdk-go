package matrix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/version"
)

// Interface consists of functions of different functionalities of Matrix service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetMatrix will receive a list of points as sources and a list of points as targets and returns a matrix of eta predictions from all sources to all targets.
	// Will return error if sources or targets are empty.
	GetMatrix(sources []Point, targets []Point, options CallOptions) (Output, error)
	// GetMatrixWithContext s like GetMatrix, but with context.Context support
	GetMatrixWithContext(ctx context.Context, sources []Point, targets []Point, options CallOptions) (Output, error)
}

type Version string

const (
	V1 Version = "v1"
	V2 Version = "v2"

	NoTrafficQueryParameter = "no_traffic"
	EngineQueryParameter    = "engine"
	JSONInputQueryParam     = "json"
)

// Client is the main implementation of Interface for area-gateways service
type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
	tracerName string
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

// GetMatrix will receive a list of points as sources and a list of points as targets and returns a matrix of eta predictions from all sources to all targets.
// Will return error if sources or targets are empty.
func (c *Client) GetMatrix(sources []Point, targets []Point, options CallOptions) (Output, error) {
	return c.GetMatrixWithContext(context.Background(), sources, targets, options)
}

// GetMatrixWithContext s like GetMatrix, but with context.Context support
func (c *Client) GetMatrixWithContext(ctx context.Context, sources []Point, targets []Point, options CallOptions) (Output, error) {
	if ctx == nil {
		return Output{}, fmt.Errorf("smapp matrix: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-matrix")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	if len(sources) == 0 || len(targets) == 0 {
		return Output{}, fmt.Errorf("smapp matrix: both sources and targets should not be empty")
	}

	params := url.Values{}
	if options.UseNoTraffic {
		params.Set(NoTrafficQueryParameter, strconv.FormatBool(options.NoTraffic))
	}

	if options.EngineStr != "" {
		params.Set(EngineQueryParameter, options.EngineStr)
	} else {
		params.Set(EngineQueryParameter, options.Engine.String())
	}

	input := Input{Sources: sources, Targets: targets}
	var (
		req *http.Request
		err error
	)

	if options.UsePost {
		postInput := PostInput{input}
		// ---------- HTTP POST ----------
		body, err := json.Marshal(postInput)
		if err != nil {
			reqInitSpan.RecordError(err)
			reqInitSpan.End()
			return Output{}, fmt.Errorf("smapp matrix: could not marshal input data: %w", err)
		}

		req, err = http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			c.url,
			bytes.NewReader(body),
		)
		if err != nil {
			reqInitSpan.RecordError(err)
			reqInitSpan.End()
			return Output{}, fmt.Errorf("smapp matrix: could not create POST request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

	} else {
		// ---------- HTTP GET (legacy) ----------
		jsonData, err := json.Marshal(input)
		if err != nil {
			reqInitSpan.RecordError(err)
			reqInitSpan.End()
			return Output{}, fmt.Errorf("smapp matrix: could not marshal input data: %w", err)
		}
		params.Set(JSONInputQueryParam, string(jsonData))

		req, err = http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
		if err != nil {
			reqInitSpan.RecordError(err)
			reqInitSpan.End()
			return Output{}, fmt.Errorf("smapp matrix: could not create GET request: %w", err)
		}
	}

	// ---- API-key handling ----
	switch c.cfg.APIKeySource {
	case config.HeaderSource:
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	case config.QueryParamSource:
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	default:
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: invalid api key source: %s", c.cfg.APIKeySource)
	}

	// apply query string (for both GET and POST paths)
	req.URL.RawQuery = params.Encode()

	// extra headers
	for k, v := range options.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	reqInitSpan.End()

	// ---- perform request ----
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Output{}, fmt.Errorf("smapp matrix: request failed: %w", err)
	}

	var respSpan trace.Span
	_, respSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		respSpan.SetStatus(codes.Error, "non 200 status code")
		respSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: non 200 status: %d", resp.StatusCode)
	}

	var out Output
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		respSpan.RecordError(err)
		respSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: could not decode response: %w", err)
	}

	respSpan.End()
	return out, nil
}

// NewMatrixClient is the constructor of Matrix client.
func NewMatrixClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg:       cfg,
		httpClient: http.Client{
			Timeout:   timeout,
			Transport: http.DefaultTransport,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.url == "" {
		client.url = getMatrixDefaultURL(cfg, version)
	}

	return client, nil
}

func getMatrixDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	if version != V1{
		// New upstream layout: {base}/api/{version}/matrix
		return fmt.Sprintf("%s/api/%s/matrix", baseURL, version)
	} else {
		// Legacy layout: {base}/matrix/{version}
		return fmt.Sprintf("%s/matrix/%s", baseURL, version)
	}
}
