package config

import "strings"

type Option func(config *Config)

func WithRegion(region string) Option{
	return func(config *Config) {
		config.Region = region
	}
}

func WithAPIKey(apiKey string) Option {
	return func(config *Config) {
		config.APIKey = apiKey
	}
}

func WithAPIBaseURL(baseURL string) Option {
	return func(config *Config) {
		config.APIBaseURL = baseURL
	}
}

func WithAPIKeySource(source APIKeySource) Option {
	return func(config *Config) {
		config.APIKeySource = source
	}
}

func WithAPIKeyName(keyName string) Option {
	return func(config *Config) {
		config.APIKeyName = keyName
	}
}

func WithPublicURL() Option {
	return func(config *Config) {
		if config.Region == "" {
			config.Region = DefaultRegion
		}
		config.APIBaseURL = strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", config.Region)
	}
}

func WithInternalURL() Option {
	return func(config *Config) {
		if config.Region == "" {
			config.Region = DefaultRegion
		}
		config.APIBaseURL = strings.ReplaceAll(InternalBaseURLPattern, "{REGION}", config.Region)
	}
}