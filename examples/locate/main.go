package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/locate"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}
	client, err := locate.NewLocateClient(cfg, locate.V1, time.Second*10,
		locate.WithURL("https://bifrost.apps.private.teh-2.snappcloud.io/api/v1/locate"))
	if err != nil {
		panic(err)
	}

	results, err := client.LocatePoints([]locate.Point{{
		Lat: 35.70973799747619,
		Lon: 51.40869855880737,
	}}, locate.NewDefaultCallOptions(
		locate.WithHeaders(map[string]string{
			"foo": "bar",
		})))
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
