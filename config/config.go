package config

import (
	"os"
	"strings"
)

type APIKeySource string

const (
	HeaderSource     APIKeySource = "header"
	QueryParamSource APIKeySource = "query"

	DefaultHeaderAPIKeyName     = "X-Monshi-Key"
	DefaultQueryParamAPIKeyName = "monshi_key"

	InternalBaseURLPattern = "http://smapp-api.apps.inter-dc.{REGION}.snappcloud.io"
	PublicBaseURLPattern = "http://api.{REGION}.snappmaps.ir/"

	DefaultBaseURL = "http://smapp-api.apps.inter-dc.teh-1.snappcloud.io"
	DefaultRegion  = "teh-1"
)

type Config struct {
	Region       string
	APIKey       string
	APIKeySource APIKeySource
	APIKeyName   string
	APIBaseURL   string
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

func ReadFromEnvironment(opts ...Option) (*Config, error) {
	apiKey := os.Getenv("SMAPP_API_KEY")
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}
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

	err := config.setDefaults()
	if err != nil {
		return nil, err
	}

	return config, nil
}

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
