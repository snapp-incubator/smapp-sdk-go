package search

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Version string

type Interface interface {
	GetCities(options CallOptions) ([]City, error)
	GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error)
	SearchCity(input string, options CallOptions) ([]City, error)
	SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error)
	AutoComplete(input string, options CallOptions) ([]Result, error)
	AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error)
	Details(placeId string, options CallOptions) (Detail, error)
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

type Client struct {
	cfg        *config.Config
	url        string
	httpClient http.Client
}

// Force Client to implement Interface at compile time
var _ Interface = (*Client)(nil)

func (c *Client) GetCities(options CallOptions) ([]City, error) {
	return c.GetCitiesWithContext(context.Background(), options)
}

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

	if options.Language != "" {
		params.Set(Lang, string(options.Language))
	}

	if options.RequestContext != "" {
		params.Set(ReqContext, string(options.RequestContext))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp search get-cities: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

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

		if resp.Status != OKStatus {
			return nil, errors.New("smapp search get-cities: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search get-cities: non 200 status: %d", response.StatusCode)
}

func (c *Client) SearchCity(input string, options CallOptions) ([]City, error) {
	return c.SearchCityWithContext(context.Background(), input, options)
}

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

	if options.Language != "" {
		params.Set(Lang, string(options.Language))
	}

	if options.RequestContext != "" {
		params.Set(ReqContext, string(options.RequestContext))
	}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		return nil, fmt.Errorf("smapp search search-cities: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

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

		if resp.Status != OKStatus {
			return nil, errors.New("smapp search search-cities: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search search-cities: non 200 status: %d", response.StatusCode)
}

func (c *Client) AutoComplete(input string, options CallOptions) ([]Result, error) {
	return c.AutoCompleteWithContext(context.Background(), input, options)
}

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

	if options.Language != "" {
		params.Set(Lang, string(options.Language))
	}

	if options.UseUserLocation {
		locationString := fmt.Sprintf("%f,%f", options.UserLocation.Lat, options.UserLocation.Lon)
		params.Set(UserLocation, locationString)
	}

	if options.RequestContext != "" {
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

		if resp.Status != OKStatus {
			return nil, errors.New("smapp search autocomplete: status of request is not OK")
		}

		return resp.Predictions, nil
	}

	return nil, fmt.Errorf("smapp search autocomplete: non 200 status: %d", response.StatusCode)
}

func (c *Client) Details(placeId string, options CallOptions) (Detail, error) {
	return c.DetailsWithContext(context.Background(), placeId, options)
}

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

		if resp.Status != OKStatus {
			return Detail{}, errors.New("smapp search details: status of request is not OK")
		}

		return resp.Result, nil
	}

	return Detail{}, fmt.Errorf("smapp search details: non 200 status: %d", response.StatusCode)
}

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
