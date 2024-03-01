package main

import (
	"fmt"
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/eta"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}
	client, err := eta.NewETAClient(cfg, eta.V1, time.Second*10,
		eta.WithURL("https://bifrost.apps.private.teh-2.snappcloud.io/api/v1/eta"))
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
		eta.WithNoTraffic(),
	),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
