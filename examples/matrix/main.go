package main

import (
	"fmt"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/matrix"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key", 
		config.WithAPIBaseURL("https://bifrost.apps.private.okd4.teh-2.snappcloud.io"),
		config.WithAPIKeyName("X-API-Key"))
	if err != nil {
		panic(err)
	}
	client, err := matrix.NewMatrixClient(cfg, matrix.V2, time.Second*10)
	if err != nil {
		panic(err)
	}

	result, err := client.GetMatrix([]matrix.Point{
		{
			Lat: 35.7733304928583,
			Lon: 51.418322660028934,
		},
		{
			Lat: 35.72895575080859,
			Lon: 51.37228488922119,
		},
	}, []matrix.Point{
		{
			Lat: 35.70033104179786,
			Lon: 51.351492404937744,
		},
		{
			Lat: 35.73933685292328,
			Lon: 51.50890588760376,
		},
	}, matrix.NewDefaultCallOptions(
		matrix.WithHeaders(map[string]string{
			"foo": "bar",
		}),
	))

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
