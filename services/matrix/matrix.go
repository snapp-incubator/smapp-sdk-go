package matrix

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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

	NoTrafficQueryParameter = "no_traffic"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: could not create request. err: %s", err.Error())
	}

	params := url.Values{}
	if options.UseNoTraffic {
		params.Set(NoTrafficQueryParameter, "true")
	}

	data := Input{
		Sources: sources,
		Targets: targets,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: could not marshal input data")
	}

	params.Set(JSONInputQueryParam, string(jsonData))

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return Output{}, fmt.Errorf("smapp matrix: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()
	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Output{}, fmt.Errorf("smapp matrix: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		var result Output
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return Output{}, fmt.Errorf("smapp matrix: could not serialize response due to: %s", err.Error())
		}
		responseSpan.End()
		return result, nil
	}
	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return Output{}, fmt.Errorf("smapp matrix: non 200 status: %d", response.StatusCode)
}

// NewMatrixClient is the constructor of Matrix client.
func NewMatrixClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getMatrixDefaultURL(cfg, version),
		httpClient: http.Client{
			Timeout:   timeout,
			Transport: http.DefaultTransport,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func getMatrixDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/matrix/%s", baseURL, version)
}
