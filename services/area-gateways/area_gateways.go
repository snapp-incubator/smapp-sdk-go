package area_gateways

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/version"
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
	params := url.Values{}

	point := Point{
		Lat: lat,
		Lon: lon,
	}

	err := point.Validate()
	if err != nil {
		return Area{}, fmt.Errorf("smapp area-gateways: input lat and lon are invalid due to: %s", err.Error())
	}

	body, err := json.Marshal(&point)
	if err != nil {
		return Area{}, fmt.Errorf("smapp area-gateways: could not marshal request body due to: %s", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, bytes.NewBuffer(body))
	if err != nil {
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
		return Area{}, fmt.Errorf("smapp area-gateways: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Area{}, fmt.Errorf("smapp area-gateways: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := Area{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return Area{}, fmt.Errorf("smapp area-gateways: could not serialize response due to: %s", err.Error())
		}

		return resp, nil
	}

	return Area{}, fmt.Errorf("smapp area-gateways: non 200 status: %d", response.StatusCode)
}

// NewAreaGatewaysClient is the constructor of area-gateways client.
func NewAreaGatewaysClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getAreaGatewaysDefaultURL(cfg, version),
		httpClient: http.Client{
			Timeout: timeout,
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
