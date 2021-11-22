package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/area-gateways"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	areaGatewaysClient, err := area_gateways.NewAreaGatewaysClient(cfg, area_gateways.V1, time.Second,
		area_gateways.WithURL("https://area-gateways.apps.private.teh-2.snappcloud.io/gateways"),
	)
	if err != nil {
		panic(err)
	}

	area, err := areaGatewaysClient.GetGateways(35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(area)
}
