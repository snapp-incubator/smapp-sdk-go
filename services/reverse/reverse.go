package reverse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Interface consists of functions of different functionalities of a reverse geocode service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetComponents receives `lat`,`lon` as a location and CallOptions and returns Component s of address of location given.
	GetComponents(lat, lon float64, options CallOptions) ([]Component, error)
	// GetDisplayName receives `lat`,`lon` as a location and CallOptions and returns a string as address of given location.
	GetDisplayName(lat, lon float64, options CallOptions) (string, error)
	// GetComponentsWithContext is like GetComponents, but with context.Context support.
	GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error)
	// GetDisplayNameWithContext is like GetDisplayName, but with context.Context support.
	GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error)
	// GetFrequent receives `lat`, `lon` as a location and CallOptions and returns FrequentAddress for the given location.
	GetFrequent(lat, lon float64, options CallOptions) (FrequentAddress, error)
	// GetFrequentWithContext is like GetFrequent, but with context.Context support
	GetFrequentWithContext(ctx context.Context, lat, lon float64, options CallOptions) (FrequentAddress, error)
	GetBatch(request BatchReverseRequest) ([]Result, error)
	GetBatchWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error)
	GetBatchDisplayName(request BatchReverseRequest) ([]ResultWithDisplayName, error)
	GetBatchDisplayNameWithContext(ctx context.Context, request BatchReverseRequest) ([]ResultWithDisplayName, error)

	GetStructuralResult(lat, lon float64, options CallOptions) (*StructuralComponent, error)
	GetStructuralResultWithContext(ctx context.Context, lat, lon float64, options CallOptions) (*StructuralComponent, error)
	GetBatchStructuralResults(request BatchReverseRequest) ([]StructuralResult, error)
	GetBatchStructuralResultsWithContext(ctx context.Context, request BatchReverseRequest) ([]StructuralResult, error)
}

type Version string

const (
	V1 Version = "v1"

	Lat       = "lat"
	Lon       = "lon"
	Lang      = "language"
	ZoomLevel = "zoom"
	Type      = "type"
	Display   = "display"

	OKStatus    = "OK"
	ErrorStatus = "ERROR"
)

// Client is the main implementation of Interface for reverse service
type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
	tracerName string
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

// GetComponents receives `lat`,`lon` as a location and CallOptions and returns Component s of address of location given.
func (c *Client) GetComponents(lat, lon float64, options CallOptions) ([]Component, error) {
	return c.GetComponentsWithContext(context.Background(), lat, lon, options)
}

// GetDisplayName receives `lat`,`lon“ as a location and CallOptions and returns a string as address of given location.
func (c *Client) GetDisplayName(lat, lon float64, options CallOptions) (string, error) {
	return c.GetDisplayNameWithContext(context.Background(), lat, lon, options)
}

// GetComponentsWithContext is like GetComponents, but with context.Context support.
func (c *Client) GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error) {
	if ctx == nil {
		return nil, fmt.Errorf("smapp reverse geo-code: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-address-components")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return nil, errors.New("smapp reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Lat, fmt.Sprintf("%f", lat))
	params.Set(Lon, fmt.Sprintf("%f", lon))
	if options.UseLanguage {
		params.Set(Lang, string(options.Language))
	}

	if options.UseZoomLevel {
		params.Set(ZoomLevel, strconv.Itoa(options.ZoomLevel))
	}

	if options.UseResponseType {
		params.Set(Type, string(options.ResponseType))
	}
	params.Set(Display, "false")

	switch c.cfg.APIKeySource {
	case config.HeaderSource:
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	case config.QueryParamSource:
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	default:
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return nil, fmt.Errorf("smapp reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()

	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status string `json:"status"`
			Result struct {
				Components []Component `json:"components"`
			} `json:"result"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return nil, fmt.Errorf("smapp reverse geo-code: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			responseSpan.SetStatus(codes.Error, "status not OK")
			responseSpan.End()
			return nil, errors.New("smapp reverse geo-code: status of request is not OK")
		}

		responseSpan.End()
		return resp.Result.Components, nil
	}

	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return nil, fmt.Errorf("smapp reverse geo-code: non 200 status: %d", response.StatusCode)
}

// GetDisplayNameWithContext is like GetDisplayName, but with context.Context support.
func (c *Client) GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error) {
	if ctx == nil {
		return "", fmt.Errorf("smapp reverse geo-code: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-display-name-address")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return "", errors.New("smapp reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Lat, fmt.Sprintf("%f", lat))
	params.Set(Lon, fmt.Sprintf("%f", lon))

	if options.UseLanguage {
		params.Set(Lang, string(options.Language))
	}

	if options.UseZoomLevel {
		params.Set(ZoomLevel, strconv.Itoa(options.ZoomLevel))
	}

	if options.UseResponseType {
		params.Set(Type, string(options.ResponseType))
	}

	params.Set(Display, "true")

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.End()
		return "", fmt.Errorf("smapp reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()

	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("smapp reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status string `json:"status"`
			Result struct {
				DisplayName string `json:"displayName"`
			} `json:"result"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return "", fmt.Errorf("smapp reverse geo-code: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			responseSpan.RecordError(err)
			responseSpan.End()
			return "", errors.New("smapp reverse geo-code: status of request is not OK")
		}

		responseSpan.End()
		return resp.Result.DisplayName, nil
	}

	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return "", fmt.Errorf("smapp reverse geo-code: non 200 status: %d", response.StatusCode)
}

// GetFrequent receives `lat`, `lon` as a location and CallOptions and returns FrequentAddress for the given location.
func (c *Client) GetFrequent(lat, lon float64, options CallOptions) (FrequentAddress, error) {
	return c.GetFrequentWithContext(context.Background(), lat, lon, options)
}

// GetFrequentWithContext is like GetFrequent, but with context.Context support
func (c *Client) GetFrequentWithContext(ctx context.Context, lat, lon float64, options CallOptions) (FrequentAddress, error) {
	if ctx == nil {
		return FrequentAddress{}, fmt.Errorf("smapp reverse geo-code: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-frequent-address")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return FrequentAddress{}, errors.New("smapp reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Lat, fmt.Sprintf("%f", lat))
	params.Set(Lon, fmt.Sprintf("%f", lon))

	if options.UseZoomLevel {
		params.Set(ZoomLevel, strconv.Itoa(options.ZoomLevel))
	}

	params.Set(Type, string(Frequent))

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.End()
		return FrequentAddress{}, fmt.Errorf("smapp reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()

	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return FrequentAddress{}, fmt.Errorf("smapp reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := FrequentAddress{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return FrequentAddress{}, fmt.Errorf("smapp reverse geo-code: could not serialize response due to: %s", err.Error())
		}

		responseSpan.End()
		return resp, nil
	}
	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return FrequentAddress{}, fmt.Errorf("smapp reverse geo-code: non 200 status: %d", response.StatusCode)
}

// NewReverseClient is the constructor of reverse geocode client.
func NewReverseClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getReverseDefaultURL(cfg, version),
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

func getReverseDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/reverse/%s", baseURL, version)
}

func (c *Client) GetStructuralResult(lat, lon float64, options CallOptions) (*StructuralComponent, error) {
	return c.GetStructuralResultWithContext(context.Background(), lat, lon, options)
}

func (c *Client) GetStructuralResultWithContext(ctx context.Context, lat, lon float64, options CallOptions) (*StructuralComponent, error) {
	components, err := c.GetComponentsWithContext(ctx, lat, lon, options)
	if err != nil {
		return nil, err
	}
	response := c.convertComponentIntoStructureModel(components)
	return response, nil
}

func (c *Client) convertComponentIntoStructureModel(components []Component) *StructuralComponent {
	response := &StructuralComponent{}
	for _, component := range components {
		if _, ok := convertReverseTypes[component.Type]; ok {
			switch component.Type {
			case province:
				response.Province = component.Name
				break
			case city:
				response.City = component.Name
				break
			case county:
				response.County = component.Name
				break
			case town:
				response.Town = component.Name
				break
			case village:
				response.Village = component.Name
				break
			case neighbourhood:
				response.Neighbourhood = component.Name
				break
			case suburb:
				response.Suburb = component.Name
				break
			case locality:
				response.Locality = component.Name
				break
			case primary:
				response.Primary = component.Name
				break
			case secondary:
				response.Secondary = component.Name
				break
			case residential:
				response.Residential = component.Name
				break
			case poi:
				response.POI = component.Name
				break
			}
		} else {
			response.ClosedWay = component.Name
		}
	}
	return response
}
