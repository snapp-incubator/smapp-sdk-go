package config

import (
	"os"
	"strings"
)

type APIKeySource string

const (
	HeaderSource     APIKeySource = "header"
	QueryParamSource APIKeySource = "query"

	DefaultHeaderAPIKeyName     = "X-Smapp-Key"
	DefaultQueryParamAPIKeyName = "monshi_key"

	InternalBaseURLPattern = "http://smapp-api.apps.inter-dc.okd4.{REGION}.snappcloud.io"
	PublicBaseURLPattern   = "https://api.{REGION}.snappmaps.ir"

	DefaultBaseURL = "http://smapp-api.apps.inter-dc.okd4.teh-1.snappcloud.io"
	DefaultRegion  = "teh-1"
)

// Config is the struct needed for constructing service clients. it consists of common needed settings for calling different services.
type Config struct {
	// Region is the region of service to be called
	Region string
	// APIKey is the key required to authenticating to different services
	APIKey string
	// APIKeySource is for defining the source of APIKey in each request. it can be header or query params.
	APIKeySource APIKeySource
	// APIKeyName is used as key of authentication in requests.
	APIKeyName string
	// APIBaseURL is the base url of all smapp services.
	APIBaseURL string
}

func (c *Config) setDefaults() error {
	if c.APIKeySource == "" {
		c.APIKeySource = HeaderSource
	} else {
		if c.APIKeySource != HeaderSource && c.APIKeySource != QueryParamSource {
			return ErrInvalidAPIKeySource
		}
	}

	if c.APIKeyName == "" {
		if c.APIKeySource == HeaderSource {
			c.APIKeyName = DefaultHeaderAPIKeyName
		} else if c.APIKeySource == QueryParamSource {
			c.APIKeyName = DefaultQueryParamAPIKeyName
		}
	}

	if c.APIBaseURL == "" {
		if c.Region == "" {
			c.Region = DefaultRegion
		}
		c.APIBaseURL = strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", c.Region)
	}

	return nil
}

// ReadFromEnvironment helps reading a config from Environment variables. you can pass different options to override each field of config.
// This function returns an error if no APIKey is defined
// these Environment variables are used to fill a config:
// 		`SMAPP_API_KEY` for APIKey
// 		`SMAPP_API_KEY_SOURCE` for APIKeySource
// 		`SMAPP_API_KEY_NAME` for APIKeyName
// 		`SMAPP_API_REGION` for Region
// 		`SMAPP_API_BASE_URL` for APIBaseURL
func ReadFromEnvironment(opts ...Option) (*Config, error) {
	apiKey := os.Getenv("SMAPP_API_KEY")
	apiKeySource := APIKeySource(os.Getenv("SMAPP_API_KEY_SOURCE"))
	apiKeyName := os.Getenv("SMAPP_API_KEY_NAME")
	region := os.Getenv("SMAPP_API_REGION")
	baseURL := os.Getenv("SMAPP_API_BASE_URL")
	config := &Config{
		Region:       region,
		APIKey:       apiKey,
		APIKeySource: apiKeySource,
		APIKeyName:   apiKeyName,
		APIBaseURL:   baseURL,
	}

	for _, opt := range opts {
		opt(config)
	}

	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	err := config.setDefaults()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// NewDefaultConfig creates a default config. you can pass different options to override each field of config.
// This function returns an error if no apiKey is defined
// these are default values of each config field:
// 		Region: teh-1
// 		APIKeySource: header
//		APIKeyName: X-Monshi-Key
//		APIBaseURL: http://smapp-api.apps.inter-dc.teh-1.snappcloud.io
func NewDefaultConfig(apiKey string, opts ...Option) (*Config, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	config := &Config{
		Region:       DefaultRegion,
		APIKey:       apiKey,
		APIKeySource: HeaderSource,
		APIKeyName:   DefaultHeaderAPIKeyName,
		APIBaseURL:   DefaultBaseURL,
	}

	for _, opt := range opts {
		opt(config)
	}

	err := config.setDefaults()
	if err != nil {
		return nil, err
	}

	return config, nil
}
