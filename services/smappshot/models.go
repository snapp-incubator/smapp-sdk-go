package smappshot

import "time"

// Version represents the SmappShot API version.
type Version string

const (
	V1 Version = "v1"
	V2 Version = "v2"
)

// Language represents supported map languages.
type Language string

const (
	LanguageFarsi   Language = "fa"
	LanguageEnglish Language = "en"
	LanguageArabic  Language = "ar"
	LanguageKurdish Language = "ku"
)

// MarkerType represents the type of map marker rendered on ride photos.
type MarkerType int

const (
	MarkerTypeRideHistory   MarkerType = 0
	MarkerTypeLocationShare MarkerType = 1
)

// Location is a geographic coordinate. Longitude comes first, matching the SmappShot wire format.
type Location struct {
	Lon float64 // longitude
	Lat float64 // latitude
}

// SigningConfig holds HMAC-SHA256 URL signing parameters.
type SigningConfig struct {
	Secret         string
	ExpiryDuration time.Duration
}
