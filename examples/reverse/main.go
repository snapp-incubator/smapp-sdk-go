package main

import (
	"fmt"
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/reverse"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	reverseClient, err := reverse.NewReverseClient(cfg, reverse.V1, time.Second,
		reverse.WithURL("https://new-url.com"), // This is optional
	)
	if err != nil {
		panic(err)
	}

	displayName, err := reverseClient.GetDisplayName(35.0123, 53.12312, reverse.NewDefaultCallOptions(
		reverse.WithZoomLevel(17),
		reverse.WithEnglishLanguage(),
	))
	if err != nil {
		panic(err)
	}

	fmt.Println(displayName)
}
