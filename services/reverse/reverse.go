package reverse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Interface interface {
	GetComponents(lat, lon float64, options CallOptions) ([]Component, error)
	GetDisplayName(lat, lon float64, options CallOptions) (string, error)
	GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error)
	GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error)
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

type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

func (c *Client) GetComponents(lat, lon float64, options CallOptions) ([]Component, error) {
	return c.GetComponentsWithContext(context.Background(), lat, lon, options)
}

func (c *Client) GetDisplayName(lat, lon float64, options CallOptions) (string, error) {
	return c.GetDisplayNameWithContext(context.Background(), lat, lon, options)
}

func (c *Client) GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		return nil, errors.New("smapp reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Lat, fmt.Sprintf("%f", lat))
	params.Set(Lon, fmt.Sprintf("%f", lon))
	params.Set(Lang, string(options.Language))
	params.Set(ZoomLevel, string(rune(options.ZoomLevel)))
	params.Set(Type, string(options.ResponseType))
	params.Set(Display, "false")

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status string `json:"status"`
			Result struct {
				Components []Component `json:"components"`
			} `json:"result"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, fmt.Errorf("smapp reverse geo-code: could not serialize response due to: %s", err.Error())
		}

		if resp.Status != OKStatus {
			return nil, errors.New("smapp reverse geo-code: status of request is not OK")
		}

		return resp.Result.Components, nil
	}

	return nil, fmt.Errorf("smapp reverse geo-code: non 200 status: %d", response.StatusCode)
}

func (c *Client) GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		return "", errors.New("smapp reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Lat, fmt.Sprintf("%f", lat))
	params.Set(Lon, fmt.Sprintf("%f", lon))
	params.Set(Lang, string(options.Language))
	params.Set(ZoomLevel, string(rune(options.ZoomLevel)))
	params.Set(Type, string(options.ResponseType))
	params.Set(Display, "true")

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return "", fmt.Errorf("smapp reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("smapp reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status string `json:"status"`
			Result struct {
				DisplayName string `json:"displayName"`
			} `json:"result"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return "", fmt.Errorf("smapp reverse geo-code: could not serialize response due to: %s", err.Error())
		}

		if resp.Status != OKStatus {
			return "", errors.New("smapp reverse geo-code: status of request is not OK")
		}

		return resp.Result.DisplayName, nil
	}

	return "", fmt.Errorf("smapp reverse geo-code: non 200 status: %d", response.StatusCode)
}

func NewReverseClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getReverseDefaultURL(cfg, version),
		httpClient: http.Client{
			Timeout: timeout,
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
