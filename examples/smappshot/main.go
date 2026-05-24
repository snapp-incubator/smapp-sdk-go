package main

import (
	"fmt"
	"log"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/services/smappshot"
)

func main() {
	cfg := smappshot.SigningConfig{
		Secret:         "my-secret", // replace with real secret
		ExpiryDuration: 10 * time.Minute,
	}

	rideURL, err := smappshot.NewRideRequestBuilder("https://smappshot.snappmaps.ir", cfg, smappshot.V1).
		WithOrigin(smappshot.Location{Lon: 51.338, Lat: 35.699}).
		WithDestinations([]smappshot.Location{
			{Lon: 51.400, Lat: 35.720},
		}).
		WithLanguage(smappshot.LanguageFarsi).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ride URL:")
	fmt.Println(rideURL)
	fmt.Println()

	previewURL, err := smappshot.NewPreviewRequestBuilder("https://smappshot.snappmaps.ir", cfg, smappshot.V1).
		WithCenter(smappshot.Location{Lon: 51.338, Lat: 35.699}).
		WithZoom(14).
		WithLanguage(smappshot.LanguageFarsi).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Preview URL:")
	fmt.Println(previewURL)
}
