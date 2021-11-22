package locate

import (
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

// Interface consists of functions of different functionalities of a locate service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// LocatePoints receives a list of Point s and returns a list with same length with located Point s
	LocatePoints(points []Point, options CallOptions) ([]Result, error)
	// LocatePointsWithContext is like LocatePoints, but with context.Context support
	LocatePointsWithContext(ctx context.Context, points []Point, options CallOptions) ([]Result, error)
}

type Version string

const (
	V1 Version = "v1"

	JSONInputQueryParam = "json"
)

// Client is the main implementation of Interface for locate service
type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

// LocatePoints receives a list of Point s and returns a list with same length with located Point s
func (c *Client) LocatePoints(points []Point, options CallOptions) ([]Result, error) {
	return c.LocatePointsWithContext(context.Background(), points, options)
}

// LocatePointsWithContext is like LocatePoints, but with context.Context support
func (c *Client) LocatePointsWithContext(ctx context.Context, points []Point, options CallOptions) ([]Result, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		return nil, fmt.Errorf("smapp locate: could not create request. err: %s", err.Error())
	}

	params := url.Values{}
	type ReqData struct {
		Locations []Point `json:"locations"`
	}
	data := ReqData{Locations: points}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("smapp locate: could not marshal input data")
	}

	params.Set(JSONInputQueryParam, string(jsonData))

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp locate: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp locate: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		var result []Result
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("smapp locate: could not serialize response due to: %s", err.Error())
		}

		return result, nil
	}

	return nil, fmt.Errorf("smapp locate: non 200 status: %d", response.StatusCode)
}

// NewLocateClient is the constructor of locate client.
func NewLocateClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getLocateDefaultURL(cfg, version),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func getLocateDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/locate/%s", baseURL, version)
}
