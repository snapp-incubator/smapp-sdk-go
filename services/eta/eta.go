package eta

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Interface consists of functions of different functionalities of ETA service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetETA will receive a list of point with minimum length of 2 and returns ETA.
	// Will return error if less than 2 points are passed.
	GetETA(points []Point, options CallOptions) (ETA, error)
	// GetETAWithContext s like GetETA, but with context.Context support
	GetETAWithContext(ctx context.Context, points []Point, options CallOptions) (ETA, error)
}

type Version string

const (
	V1 Version = "v1"

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
	if ctx == nil {
		return ETA{}, fmt.Errorf("smapp eta: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-eta")
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

	type ReqData struct {
		Locations         []Point `json:"locations"`
		DepartureDateTime string  `json:"departure_date_time,omitempty"`
	}
	data := ReqData{Locations: points}

	if options.UseDepartureDateTime {
		data.DepartureDateTime = options.DepartureDateTime
	}

	params.Set(EngineQueryParameter, options.Engine.String())

	jsonData, err := json.Marshal(data)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return ETA{}, fmt.Errorf("smapp eta: could not marshal input data")
	}

	params.Set(JSONInputQueryParam, string(jsonData))

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
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

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
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
		url: getETADefaultURL(cfg, version),
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

func getETADefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/eta/%s", baseURL, version)
}
