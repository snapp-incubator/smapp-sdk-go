package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/services/smappshot"
)

func main() {
	secret := os.Getenv("SMAPPSHOT_SIGNING_SECRET")

	// Ride — Baghdad, Iraq
	rideURL, err := smappshot.NewRideRequestBuilder("http://localhost:8080", secret, smappshot.V1).
		WithOrigin(smappshot.Location{Lat: 33.3152, Lon: 44.3661}).
		WithDestinations([]smappshot.Location{
			{Lat: 33.3600, Lon: 44.4000},
		}).
		WithLanguage(smappshot.LanguageEnglish).
		WithExpiry(10 * time.Minute).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Ride (Baghdad) ===")
	fmt.Printf("curl -o /tmp/ride.png -w \"%%{http_code}\\n\" \"%s\"\n\n", rideURL)

	// Preview — Beirut, Lebanon
	previewURL, err := smappshot.NewPreviewRequestBuilder("http://localhost:8080", secret, smappshot.V1).
		WithCenter(smappshot.Location{Lat: 33.8938, Lon: 35.5018}).
		WithZoom(14).
		WithLanguage(smappshot.LanguageEnglish).
		WithExpiry(10 * time.Minute).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Preview (Beirut) ===")
	fmt.Printf("curl -o /tmp/preview.png -w \"%%{http_code}\\n\" \"%s\"\n", previewURL)
}
