package main

import (
	"fmt"
	"log"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/services/smappshot"
)

func main() {
	secret := "my-secret" // replace with real secret

	rideURL, err := smappshot.NewRideRequestBuilder("https://smappshot.snappmaps.ir", secret, smappshot.V1).
		WithOrigin(smappshot.Location{Lon: 51.338, Lat: 35.699}).
		WithDestinations([]smappshot.Location{
			{Lon: 51.400, Lat: 35.720},
		}).
		WithLanguage(smappshot.LanguageFarsi).
		WithExpiry(10 * time.Minute).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ride URL:")
	fmt.Println(rideURL)
	fmt.Println()

	previewURL, err := smappshot.NewPreviewRequestBuilder("https://smappshot.snappmaps.ir", secret, smappshot.V1).
		WithCenter(smappshot.Location{Lon: 51.338, Lat: 35.699}).
		WithZoom(14).
		WithLanguage(smappshot.LanguageFarsi).
		WithExpiry(10 * time.Minute).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Preview URL:")
	fmt.Println(previewURL)
}
