package config

import (
	"strings"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	t.Run("empty_api_key", func(t *testing.T) {
		_, err := NewDefaultConfig("")
		if err == nil {
			t.Fatal("NewDefaultConfig err should not be nil")
		}
	})

	t.Run("without_options", func(t *testing.T) {
		c, err := NewDefaultConfig("foo")
		if err != nil {
			t.Fatalf("NewDefaultConfig should not return error: %s", err.Error())
		}
		if c.APIKey != "foo" {
			t.Fatalf("config.APIKey should be foo but it is: %s", c.APIKey)
		}
		if c.APIKeySource != HeaderSource {
			t.Fatalf("config.APIKeySource should be %s but it is: %s", HeaderSource, c.APIKeySource)
		}
		if c.APIKeyName != DefaultHeaderAPIKeyName {
			t.Fatalf("config.APIKeyName should be %s but it is: %s", DefaultHeaderAPIKeyName, c.APIKeyName)
		}
		if c.Region != DefaultRegion {
			t.Fatalf("config.Region should be %s but it is: %s", DefaultRegion, c.Region)
		}
		baseUrlValue := strings.ReplaceAll(InternalBaseURLPattern, "{REGION}", c.Region)
		if c.APIBaseURL != baseUrlValue {
			t.Fatalf("config.APIBaseURL should be %s but it is: %s", baseUrlValue, c.APIBaseURL)
		}
	})

	t.Run("with_options", func(t *testing.T) {
		t.Run("without_error", func(t *testing.T) {
			c, err := NewDefaultConfig("foo",
				WithAPIKeyName("X-Monshi-New-Key"))
			if err != nil {
				t.Fatalf("NewDefaultConfig should not return error: %s", err.Error())
			}
			if c.APIKey != "foo" {
				t.Fatalf("config.APIKey should be foo but it is: %s", c.APIKey)
			}
			if c.APIKeySource != HeaderSource {
				t.Fatalf("config.APIKeySource should be %s but it is: %s", HeaderSource, c.APIKeySource)
			}
			if c.APIKeyName != "X-Monshi-New-Key" {
				t.Fatalf("config.APIKeyName should be %s but it is: %s", "X-Monshi-New-Key", c.APIKeyName)
			}
			if c.Region != DefaultRegion {
				t.Fatalf("config.Region should be %s but it is: %s", DefaultRegion, c.Region)
			}
			baseUrlValue := strings.ReplaceAll(InternalBaseURLPattern, "{REGION}", c.Region)
			if c.APIBaseURL != baseUrlValue {
				t.Fatalf("config.APIBaseURL should be %s but it is: %s", baseUrlValue, c.APIBaseURL)
			}
		})

		t.Run("with_error", func(t *testing.T) {
			_, err := NewDefaultConfig("foo",
				WithAPIKeySource("test"))
			if err == nil {
				t.Fatalf("NewDefaultConfig should not be nill")
			}
		})
	})
}

func TestReadFromEnvironment(t *testing.T) {
	t.Run("empty_api_key", func(t *testing.T) {
		_, err := ReadFromEnvironment()
		if err == nil {
			t.Fatal("ReadFromEnvironment err should not be nil")
		}
	})

	t.Run("without_options", func(t *testing.T) {
		t.Setenv("SMAPP_API_KEY", "foo")
		t.Setenv("SMAPP_API_KEY_SOURCE", "header")
		t.Setenv("SMAPP_API_KEY_NAME", "X-Monshi")
		t.Setenv("SMAPP_API_REGION", "teh-2")
		c, err := ReadFromEnvironment()
		if err != nil {
			t.Fatalf("ReadFromEnvironment should not return error: %s", err.Error())
		}
		if c.APIKey != "foo" {
			t.Fatalf("config.APIKey should be foo but it is: %s", c.APIKey)
		}
		if c.APIKeySource != HeaderSource {
			t.Fatalf("config.APIKeySource should be %s but it is: %s", HeaderSource, c.APIKeySource)
		}
		if c.APIKeyName != "X-Monshi" {
			t.Fatalf("config.APIKeyName should be %s but it is: %s", "X-Monshi", c.APIKeyName)
		}
		if c.Region != "teh-2" {
			t.Fatalf("config.Region should be %s but it is: %s", "teh-2", c.Region)
		}
		baseUrlValue := strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", c.Region)
		if c.APIBaseURL != baseUrlValue {
			t.Fatalf("config.APIBaseURL should be %s but it is: %s", baseUrlValue, c.APIBaseURL)
		}
	})

	t.Run("with_options", func(t *testing.T) {
		t.Run("with_error", func(t *testing.T) {
			t.Setenv("SMAPP_API_KEY", "foo")
			t.Setenv("SMAPP_API_KEY_SOURCE", "header")
			t.Setenv("SMAPP_API_KEY_NAME", "X-Monshi")
			t.Setenv("SMAPP_API_REGION", "teh-2")
			_, err := ReadFromEnvironment(
				WithAPIKeySource("bar"))
			if err == nil {
				t.Fatalf("ReadFromEnvironment should return error")
			}
		})

		t.Run("without_error", func(t *testing.T) {
			t.Setenv("SMAPP_API_KEY", "foo")
			t.Setenv("SMAPP_API_KEY_SOURCE", "header")
			t.Setenv("SMAPP_API_KEY_NAME", "X-Monshi")
			t.Setenv("SMAPP_API_REGION", "teh-2")
			c, err := ReadFromEnvironment(
				WithAPIKey("bar"))
			if err != nil {
				t.Fatalf("ReadFromEnvironment should not return error: %s", err.Error())
			}
			if c.APIKey != "bar" {
				t.Fatalf("config.APIKey should be bar but it is: %s", c.APIKey)
			}
			if c.APIKeySource != HeaderSource {
				t.Fatalf("config.APIKeySource should be %s but it is: %s", HeaderSource, c.APIKeySource)
			}
			if c.APIKeyName != "X-Monshi" {
				t.Fatalf("config.APIKeyName should be %s but it is: %s", "X-Monshi", c.APIKeyName)
			}
			if c.Region != "teh-2" {
				t.Fatalf("config.Region should be %s but it is: %s", "teh-2", c.Region)
			}
			baseUrlValue := strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", c.Region)
			if c.APIBaseURL != baseUrlValue {
				t.Fatalf("config.APIBaseURL should be %s but it is: %s", baseUrlValue, c.APIBaseURL)
			}
		})
	})
}

func TestConfig_setDefaults(t *testing.T) {
	t.Run("empty_all", func(t *testing.T) {
		c := &Config{}
		err := c.setDefaults()
		if err != nil {
			t.Fatalf("setDefaults should not return error: %s", err.Error())
		}

		if c.APIKeySource != HeaderSource {
			t.Fatalf("APIKeySource should be %s but t is %s", HeaderSource, c.APIKeySource)
		}

		if c.APIKeyName != DefaultHeaderAPIKeyName {
			t.Fatalf("APIKeyName should be %s but t is %s", DefaultHeaderAPIKeyName, c.APIKeyName)
		}

		if c.Region != DefaultRegion {
			t.Fatalf("Region should be %s but t is %s", DefaultRegion, c.Region)
		}

		if c.APIBaseURL != strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", c.Region) {
			t.Fatalf("APIBaseURL should not be %s", c.APIBaseURL)
		}
	})

	t.Run("different_api_key_source", func(t *testing.T) {
		c := &Config{
			APIKeySource: QueryParamSource,
		}
		err := c.setDefaults()
		if err != nil {
			t.Fatalf("setDefaults should not return error: %s", err.Error())
		}

		if c.APIKeySource != QueryParamSource {
			t.Fatalf("APIKeySource should be %s but t is %s", QueryParamSource, c.APIKeySource)
		}

		if c.APIKeyName != DefaultQueryParamAPIKeyName {
			t.Fatalf("APIKeyName should be %s but t is %s", DefaultQueryParamAPIKeyName, c.APIKeyName)
		}

		if c.Region != DefaultRegion {
			t.Fatalf("Region should be %s but t is %s", DefaultRegion, c.Region)
		}

		if c.APIBaseURL != strings.ReplaceAll(PublicBaseURLPattern, "{REGION}", c.Region) {
			t.Fatalf("APIBaseURL should not be %s", c.APIBaseURL)
		}
	})
}
