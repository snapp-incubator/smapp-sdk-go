package config

import (
	"strings"
	"testing"
)

func TestWithAPIBaseURL(t *testing.T) {
	c, err := NewDefaultConfig("foo", WithAPIBaseURL("https://google.com"))
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	if c.APIBaseURL != "https://google.com" {
		t.Fatalf("APIBaseURL should be %s but it is %s", "https://google.com", c.APIBaseURL)
	}
}

func TestWithAPIKey(t *testing.T) {
	c, err := NewDefaultConfig("foo", WithAPIKey("bar"))
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	if c.APIKey != "bar" {
		t.Fatalf("APIKey should be %s but it is %s", "bar", c.APIKey)
	}
}

func TestWithAPIKeyName(t *testing.T) {
	c, err := NewDefaultConfig("foo", WithAPIKeyName("X-Bar"))
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	if c.APIKeyName != "X-Bar" {
		t.Fatalf("APIKeyName should be %s but it is %s", "X-Bar", c.APIKeyName)
	}
}

func TestWithAPIKeySource(t *testing.T) {
	c, err := NewDefaultConfig("foo", WithAPIKeySource("query"))
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	if c.APIKeySource != "query" {
		t.Fatalf("APIKeySource should be %s but it is %s", "query", c.APIKeySource)
	}
}

func TestWithInternalURL(t *testing.T) {
	c, err := NewDefaultConfig("foo",
		WithRegion(""),
		WithInternalURL())
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	baseUrl := strings.ReplaceAll(InternalBaseURLPattern, "{REGION}", DefaultRegion)
	if c.APIBaseURL != baseUrl {
		t.Fatalf("APIBaseURL should be %s but it is %s", baseUrl, c.APIBaseURL)
	}
}

func TestWithPublicURL(t *testing.T) {
	c, err := NewDefaultConfig("foo",
		WithRegion(""),
		WithPublicURL())
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	baseUrl := strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", DefaultRegion)
	if c.APIBaseURL != baseUrl {
		t.Fatalf("APIBaseURL should be %s but it is %s", baseUrl, c.APIBaseURL)
	}
}

func TestWithRegion(t *testing.T) {
	c, err := NewDefaultConfig("foo", WithRegion("teh-2"))
	if err != nil {
		t.Fatalf("should not return error: %s", err.Error())
	}

	if c.Region != "teh-2" {
		t.Fatalf("Region should be %s but it is %s", "teh-2", c.Region)
	}
}
