package main

import (
	"fmt"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/eta"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key", 
		config.WithAPIBaseURL("https://bifrost.apps.private.okd4.teh-2.snappcloud.io"),
		config.WithAPIKeyName("X-API-Key"))
	if err != nil {
		panic(err)
	}
	client, err := eta.NewETAClient(cfg, eta.V2, time.Second*10)
	if err != nil {
		panic(err)
	}

	results, err := client.GetETA([]eta.Point{
		{
			Lat: 35.77330981921435,
			Lon: 51.41834378242493,
		},
		{
			Lat: 35.739136559226864,
			Lon: 51.510804891586304,
		},
	}, eta.NewDefaultCallOptions(
		eta.WithHeaders(map[string]string{
			"foo": "bar",
		}),
	),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
