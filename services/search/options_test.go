package search

import (
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"testing"
	"time"
)

func TestWithURL(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewSearchClient(cfg, V1, time.Second, WithURL("https://google.com"))
	if err != nil {
		t.Fatalf("could not create search client due to: %s", err.Error())
	}

	if client.url != "https://google.com" {
		t.Fatalf("client.URL should be %s but it is %s", "https://google.com", client.url)
	}
}
