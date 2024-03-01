package main

import (
	"fmt"
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/search"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	searchClient, err := search.NewSearchClient(cfg, search.V1, time.Second,
		search.WithURL("https://new-url.com"), // This is optional
	)
	if err != nil {
		panic(err)
	}

	results, err := searchClient.AutoComplete("example", search.NewDefaultCallOptions(
		search.WithCityId(1000),
		search.WithLocation(35.012, 53.1253),
	))
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
