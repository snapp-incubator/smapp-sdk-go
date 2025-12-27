package eta

import (
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

// Interface consists of functions of different functionalities of ETA service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetETA will receive a list of point with minimum length of 2 and returns ETA.
	// Will return error if less than 2 points are passed.
	GetETA(points []Point, options CallOptions) (ETA, error)
	// GetETAWithContext s like GetETA, but with context.Context support
	GetETAWithContext(ctx context.Context, points []Point, options CallOptions) (ETA, error)
	// GetETAWithInputMeta is like GetETAWithContext, but with request-level metadata support
	GetETAWithInputMeta(ctx context.Context, points []Point, options CallOptions, metadata map[string]string) (ETA, error)
}

type Version string

const (
	V1 Version = "v1"
	V2 Version = "v2"

	NoTrafficQueryParameter = "no_traffic"
	JSONInputQueryParam     = "json"
	EngineQueryParameter    = "engine"
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

// GetETA will receive a list of point with minimum length of 2 and returns ETA.
// Will return error if less than 2 points are passed.
func (c *Client) GetETA(points []Point, options CallOptions) (ETA, error) {
	return c.GetETAWithContext(context.Background(), points, options)
}

// GetETAWithContext s like GetETA, but with context.Context support
func (c *Client) GetETAWithContext(ctx context.Context, points []Point, options CallOptions) (ETA, error) {
	return c.GetETAWithInputMeta(ctx, points, options, nil)
}

// GetETAWithInputMeta is like GetETAWithContext, but with request-level metadata support
func (c *Client) GetETAWithInputMeta(ctx context.Context, points []Point, options CallOptions, metadata map[string]string) (ETA, error) {
	if ctx == nil {
		return ETA{}, fmt.Errorf("smapp eta: nil context")
	}
	// Start of parent span
	var span trace.Span
	spanName := "get-eta"
	if len(metadata) > 0 {
		spanName = "get-eta-with-input-meta"
	}
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, spanName)
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return ETA{}, fmt.Errorf("smapp eta: could not create request. err: %s", err.Error())
	}

	params := url.Values{}
	if options.UseNoTraffic {
		params.Set(NoTrafficQueryParameter, strconv.FormatBool(options.NoTraffic))
	}

	data := ETARequest{Locations: points}
	if len(metadata) > 0 {
		data.Metadata = metadata
	}

	if options.UseDepartureDateTime {
		data.DepartureDateTime = options.DepartureDateTime
	}

	if options.EngineStr != "" {
		params.Set(EngineQueryParameter, options.EngineStr)
	} else {
		params.Set(EngineQueryParameter, options.Engine.String())
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return ETA{}, fmt.Errorf("smapp eta: could not marshal input data")
	}

	params.Set(JSONInputQueryParam, string(jsonData))

	switch c.cfg.APIKeySource {
	case config.HeaderSource:
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	case config.QueryParamSource:
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	default:
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return ETA{}, fmt.Errorf("smapp eta: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()
	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return ETA{}, fmt.Errorf("smapp eta: could not make a request due to this error: %s", err.Error())
	}

	var responseSpan trace.Span
	_, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		var result ETA
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return ETA{}, fmt.Errorf("smapp eta: could not serialize response due to: %s", err.Error())
		}
		responseSpan.End()
		return result, nil
	}
	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return ETA{}, fmt.Errorf("smapp eta: non 200 status: %d", response.StatusCode)
}

// NewETAClient is the constructor of ETA client.
func NewETAClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		httpClient: http.Client{
			Timeout:   timeout,
			Transport: http.DefaultTransport,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.url == "" {
		client.url = getETADefaultURL(cfg, version)
	}

	return client, nil
}

func getETADefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	if version != V1 {
		// New upstream layout: {base}/api/{version}/eta
		return fmt.Sprintf("%s/api/%s/eta", baseURL, version)
	} else {
		// Legacy layout: {base}/eta/{version}
		return fmt.Sprintf("%s/eta/%s", baseURL, version)
	}
}
