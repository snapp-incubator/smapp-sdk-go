package config

import "errors"

var ErrEmptyAPIKey = errors.New("api key is required")
var ErrInvalidAPIKeySource = errors.New("api key source is invalid: should be header or query")
