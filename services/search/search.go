package search

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/version"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Version string

// Interface consists of functions of different functionalities of a search service. there are two implementation of this service.
// one for mocking and one for production usage.
type Interface interface {
	// GetCities receives cCallOptions and returns list of popular City s.
	GetCities(options CallOptions) ([]City, error)
	// GetCitiesWithContext is like GetCities, but with context.Context support.
	GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error)
	// SearchCity  receives an input string for search and CallOptions and returns list of City s according to input string.
	SearchCity(input string, options CallOptions) ([]City, error)
	// SearchCityWithContext is like SearchCity, but with context.Context support.
	SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error)
	// AutoComplete receives an input string and CallOptions and returns all possible Result s according to input string.
	AutoComplete(input string, options CallOptions) ([]Result, error)
	// AutoCompleteWithContext is like AutoComplete, but with context.Context support.
	AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error)
	// Details receives a `placeId` string and CallOptions and returns Details on that place id.
	Details(placeId string, options CallOptions) (Detail, error)
	// DetailsWithContext is like Details, but with context.Context support.
	DetailsWithContext(ctx context.Context, placeId string, options CallOptions) (Detail, error)
}

const (
	V1 = "v1"

	Location     = "location"
	UserLocation = "user_location"
	Lang         = "language"
	ReqContext   = "context"
	CityID       = "city_id"
	PlaceID      = "placeid"
	Input        = "input"

	OKStatus    = "OK"
	ErrorStatus = "ERROR"
)

// Client is the main implementation of Interface for search service
type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

// GetCities receives cCallOptions and returns list of popular City s.
func (c *Client) GetCities(options CallOptions) ([]City, error) {
	return c.GetCitiesWithContext(context.Background(), options)
}

// GetCitiesWithContext is like GetCities, but with context.Context support.
func (c *Client) GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error) {
	reqURL := fmt.Sprintf("%s/place/cities", c.url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, errors.New("smapp search get-cities: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	if options.UseLocation {
		locationString := fmt.Sprintf("%f,%f", options.Location.Lat, options.Location.Lon)
		params.Set(Location, locationString)
	}

	if options.UseLanguage {
		params.Set(Lang, string(options.Language))
	}

	if options.UseRequestContext {
		params.Set(ReqContext, string(options.RequestContext))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp search get-cities: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())
	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp search get-cities: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status      string `json:"status"`
			Predictions []City `json:"predictions"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, fmt.Errorf("smapp search get-cities: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			return nil, errors.New("smapp search get-cities: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search get-cities: non 200 status: %d", response.StatusCode)
}

// SearchCity  receives an input string for search and CallOptions and returns list of City s according to input string.
func (c *Client) SearchCity(input string, options CallOptions) ([]City, error) {
	return c.SearchCityWithContext(context.Background(), input, options)
}

// SearchCityWithContext is like SearchCity, but with context.Context support.
func (c *Client) SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error) {
	reqURL := fmt.Sprintf("%s/place/search/city", c.url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, errors.New("smapp search search-cities: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Input, input)

	if options.UseLocation {
		locationString := fmt.Sprintf("%f,%f", options.Location.Lat, options.Location.Lon)
		params.Set(Location, locationString)
	}

	if options.UseLanguage {
		params.Set(Lang, string(options.Language))
	}

	if options.UseRequestContext {
		params.Set(ReqContext, string(options.RequestContext))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp search search-cities: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())
	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp search search-cities: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status      string `json:"status"`
			Predictions []City `json:"predictions"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, fmt.Errorf("smapp search search-cities: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			return nil, errors.New("smapp search search-cities: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search search-cities: non 200 status: %d", response.StatusCode)
}

// AutoComplete receives an input string and CallOptions and returns all possible Result s according to input string.
func (c *Client) AutoComplete(input string, options CallOptions) ([]Result, error) {
	return c.AutoCompleteWithContext(context.Background(), input, options)
}

// AutoCompleteWithContext is like AutoComplete, but with context.Context support.
func (c *Client) AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error) {
	reqURL := fmt.Sprintf("%s/place/autocomplete/json", c.url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, errors.New("smapp search autocomplete: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(Input, input)

	if options.UseLocation {
		locationString := fmt.Sprintf("%f,%f", options.Location.Lat, options.Location.Lon)
		params.Set(Location, locationString)
	}

	if options.UseLanguage {
		params.Set(Lang, string(options.Language))
	}

	if options.UseUserLocation {
		locationString := fmt.Sprintf("%f,%f", options.UserLocation.Lat, options.UserLocation.Lon)
		params.Set(UserLocation, locationString)
	}

	if options.UseRequestContext {
		params.Set(ReqContext, string(options.RequestContext))
	}

	if options.UseCityID {
		params.Set(CityID, strconv.Itoa(options.CityID))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp search autocomplete: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())
	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp search autocomplete: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status      string   `json:"status"`
			Predictions []Result `json:"predictions"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, fmt.Errorf("smapp search autocomplete: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			return nil, errors.New("smapp search autocomplete: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search autocomplete: non 200 status: %d", response.StatusCode)
}

// Details receives a `placeId` string and CallOptions and returns Details on that place id.
func (c *Client) Details(placeId string, options CallOptions) (Detail, error) {
	return c.DetailsWithContext(context.Background(), placeId, options)
}

// DetailsWithContext is like Details, but with context.Context support.
func (c *Client) DetailsWithContext(ctx context.Context, placeId string, options CallOptions) (Detail, error) {
	reqURL := fmt.Sprintf("%s/place/details/json", c.url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return Detail{}, errors.New("smapp search details: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	params.Set(PlaceID, placeId)

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return Detail{}, fmt.Errorf("smapp search details: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())
	for key, val := range options.Headers {
		req.Header.Set(key, val)
	}

	req.URL.RawQuery = params.Encode()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Detail{}, fmt.Errorf("smapp search details: could not make a request due to this error: %s", err.Error())
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode == http.StatusOK {
		resp := struct {
			Status string `json:"status"`
			Result Detail `json:"result"`
		}{}

		err := json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return Detail{}, fmt.Errorf("smapp search details: could not serialize response due to: %s", err.Error())
		}

		if strings.ToUpper(resp.Status) != OKStatus {
			return Detail{}, fmt.Errorf("smapp search details: request status is not OK")
		}

		return resp.Result, nil
	}

	return Detail{}, fmt.Errorf("smapp search details: non 200 status: %d", response.StatusCode)
}

// NewSearchClient is the constructor of search client.
func NewSearchClient(cfg *config.Config, version Version, timeout time.Duration, opts ...ConstructorOption) (*Client, error) {
	client := &Client{
		cfg: cfg,
		url: getSearchDefaultURL(cfg, version),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func getSearchDefaultURL(cfg *config.Config, version Version) string {
	baseURL := strings.TrimRight(cfg.APIBaseURL, "/")
	return fmt.Sprintf("%s/search/%s", baseURL, version)
}
