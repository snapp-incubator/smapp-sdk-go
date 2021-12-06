package area_gateways

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Interface consists of functions of different functionalities of a area-gateways service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetGateways receives `lat`,`lon` as a location and CallOptions and returns a polygon and its Gateways.
	// It will return an Empty area if no area is found with given lat and lon.
	GetGateways(lat, lon float64, options CallOptions) (Area, error)
	// GetGatewaysWithContext is like GetGateways, but with context.Context support
	GetGatewaysWithContext(ctx context.Context, lat, lon float64, options CallOptions) (Area, error)
}

type Version string

const (
	V1 Version = "v1"

	AcceptLanguageHeader = "Accept-Language"
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

// GetGateways receives `lat`,`lon` as a location and CallOptions and returns a polygon and its Gateways.
// It will return an Empty area if no area is found with given lat and lon.
func (c *Client) GetGateways(lat, lon float64, options CallOptions) (Area, error) {
	return c.GetGatewaysWithContext(context.Background(), lat, lon, options)
}

// GetGatewaysWithContext is like GetGateways, but with context.Context support
func (c *Client) GetGatewaysWithContext(ctx context.Context, lat, lon float64, options CallOptions) (Area, error) {
	if ctx == nil {
		return Area{}, fmt.Errorf("smapp area-gateways: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-gateways")
	defer span.End()
	span.SetAttributes(
		attribute.Float64("lat", lat),
		attribute.Float64("lon", lon),
	)

	// Request initialization span start
	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	params := url.Values{}
	point := Point{
		Lat: lat,
		Lon: lon,
	}

	err := point.Validate()
	if err != nil {
		reqInitSpan.RecordError(err, trace.WithAttributes(
			attribute.Float64("lat", point.Lat),
			attribute.Float64("lon", point.Lon),
		))
		reqInitSpan.End()
		return Area{}, fmt.Errorf("smapp area-gateways: input lat and lon are invalid due to: %s", err.Error())
	}

	body, err := json.Marshal(&point)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return Area{}, fmt.Errorf("smapp area-gateways: could not marshal request body due to: %s", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, bytes.NewBuffer(body))
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return Area{}, fmt.Errorf("smapp area-gateways: could not create request. err: %s", err.Error())
	}

	if options.UseLanguage {
		req.Header.Set(AcceptLanguageHeader, string(options.Language))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return Area{}, fmt.Errorf("smapp area-gateways: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()
	// End of request initialization
	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Area{}, fmt.Errorf("smapp area-gateways: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	var responseSpan trace.Span
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	if response.StatusCode == http.StatusOK {
		resp := Area{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return Area{}, fmt.Errorf("smapp area-gateways: could not serialize response due to: %s", err.Error())
		}

		responseSpan.End()
		return resp, nil
	}
	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return Area{}, fmt.Errorf("smapp area-gateways: non 200 status: %d", response.StatusCode)
}

// NewAreaGatewaysClient is the constructor of area-gateways client.
func NewAreaGatewaysClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getAreaGatewaysDefaultURL(cfg, version),
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

func getAreaGatewaysDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/area-gateways/%s", baseURL, version)
}
