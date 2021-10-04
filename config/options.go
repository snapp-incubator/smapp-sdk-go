package config

import "strings"

// Option is a function type for overriding different fields of config struct in a fluent way.
// You can pass as many as options to config constructors as you want.
//
// Example:
// 		cfg, err := ReadFromEnvironment(
//		WithRegion("teh-2"),
//		WithAPIKey("example-api-key"),
//		WithPublicURL(),
//		)
type Option func(config *Config)

// WithRegion sets a region for the config.
//
// Example:
// 		cfg, err := ReadFromEnvironment(WithRegion("teh-2"))
func WithRegion(region string) Option {
	return func(config *Config) {
		config.Region = region
	}
}

// WithAPIKey sets the APIKey for the config. it is often used as an option in ReadFromEnvironment
//
// Example:
// 		cfg, err := ReadFromEnvironment(WithAPIKey("new-key"))
func WithAPIKey(apiKey string) Option {
	return func(config *Config) {
		config.APIKey = apiKey
	}
}

// WithAPIBaseURL sets a custom base URL for services
//
// Example:
// 		cfg, err := ReadFromEnvironment(WithAPIBaseURL("https://api.teh-2.snappmaps.ir"))
func WithAPIBaseURL(baseURL string) Option {
	return func(config *Config) {
		config.APIBaseURL = baseURL
	}
}

// WithAPIKeySource sets an APIKeySource for the config
//
// Example:
// 		cfg, err := ReadFromEnvironment(WithAPIKeySource(QueryParamSource))
func WithAPIKeySource(source APIKeySource) Option {
	return func(config *Config) {
		config.APIKeySource = source
	}
}

// WithAPIKeyName sets an APIKeyName for the config
//
// Example:
// 		cfg, err := ReadFromEnvironment(WithAPIKeyName("X-Monshi-APIKey"))
func WithAPIKeyName(keyName string) Option {
	return func(config *Config) {
		config.APIKeyName = keyName
	}
}

// WithPublicURL sets the APIBaseURL to public routes of smapp. Notice: make sure you set region before using this option. if not set `teh-1` region would be used as default region.
//
// Example:
// 		cfg, err := ReadFromEnvironment(
//		WithRegion("teh-2"),
//		WithPublicURL(),
//		) // it will set the base URL to http://api.teh-2.snappmaps.ir/
func WithPublicURL() Option {
	return func(config *Config) {
		if config.Region == "" {
			config.Region = DefaultRegion
		}
		config.APIBaseURL = strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", config.Region)
	}
}

// WithInternalURL sets the APIBaseURL to internal routes of smapp. Notice: make sure you set region before using this option. if not set `teh-1` region would be used as default region.
// Example:
// 		cfg, err := ReadFromEnvironment(
//		WithRegion("teh-2"),
//		WithInternalURL(),
//		) // it will set the base URL to http://smapp-api.apps.inter-dc.teh-2.snappcloud.io
func WithInternalURL() Option {
	return func(config *Config) {
		if config.Region == "" {
			config.Region = DefaultRegion
		}
		config.APIBaseURL = strings.ReplaceAll(InternalBaseURLPattern, "{REGION}", config.Region)
	}
}
