package smappshot

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultExpiryDuration = 10 * time.Minute

// RideRequestBuilder builds and signs a URL for the /api/{version}/photo/ride endpoint.
// Either WithHere (single-point mode) or WithOrigin+WithDestinations (route mode) must be called.
//
// Usage:
//
//	url, err := smappshot.NewRideRequestBuilder(baseURL, secret, smappshot.V1).
//	    WithOrigin(smappshot.Location{Lon: 51.338, Lat: 35.699}).
//	    WithDestinations([]smappshot.Location{{Lon: 51.400, Lat: 35.720}}).
//	    Build()
type RideRequestBuilder struct {
	baseURL        string
	secret         string
	expiryDuration time.Duration
	version        Version
	width          int
	height         int
	language       Language
	style          string
	tenant         Tenant
	here           *Location
	origin         *Location
	destinations   []Location
	markerType     MarkerType
}

// NewRideRequestBuilder creates a builder for the ride photo URL.
func NewRideRequestBuilder(baseURL string, secret string, version Version) *RideRequestBuilder {
	return &RideRequestBuilder{
		baseURL:        strings.TrimRight(baseURL, "/"),
		secret:         secret,
		expiryDuration: defaultExpiryDuration,
		version:        version,
	}
}

func (b *RideRequestBuilder) WithWidth(w int) *RideRequestBuilder {
	b.width = w
	return b
}

func (b *RideRequestBuilder) WithHeight(h int) *RideRequestBuilder {
	b.height = h
	return b
}

// WithExpiry sets how long the signed URL remains valid. Default is 10 minutes.
func (b *RideRequestBuilder) WithExpiry(d time.Duration) *RideRequestBuilder {
	b.expiryDuration = d
	return b
}

func (b *RideRequestBuilder) WithLanguage(lang Language) *RideRequestBuilder {
	b.language = lang
	return b
}

func (b *RideRequestBuilder) WithStyle(style string) *RideRequestBuilder {
	b.style = style
	return b
}

func (b *RideRequestBuilder) WithTenant(tenant Tenant) *RideRequestBuilder {
	b.tenant = tenant
	return b
}

// WithHere sets single-point mode. Clears Origin and Destinations.
func (b *RideRequestBuilder) WithHere(loc Location) *RideRequestBuilder {
	b.here = &loc
	b.origin = nil
	b.destinations = nil
	return b
}

// WithOrigin sets the route origin. Clears Here.
func (b *RideRequestBuilder) WithOrigin(loc Location) *RideRequestBuilder {
	b.origin = &loc
	b.here = nil
	return b
}

// WithDestinations sets route destinations. Clears Here.
func (b *RideRequestBuilder) WithDestinations(dests []Location) *RideRequestBuilder {
	b.destinations = dests
	b.here = nil
	return b
}

func (b *RideRequestBuilder) WithMarkerType(mt MarkerType) *RideRequestBuilder {
	b.markerType = mt
	return b
}

// Build validates, signs, and returns the ride photo URL.
func (b *RideRequestBuilder) Build() (string, error) {
	if err := b.validate(); err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("width", strconv.Itoa(defaultInt(b.width, 512)))
	params.Set("height", strconv.Itoa(defaultInt(b.height, 285)))

	if b.here != nil {
		params.Set("here", formatLocation(*b.here))
	} else {
		params.Set("origin", formatLocation(*b.origin))
		dests := make([]string, len(b.destinations))
		for i, d := range b.destinations {
			dests[i] = formatLocation(d)
		}
		params.Set("destinations", strings.Join(dests, ";"))
	}

	params.Set("language", string(defaultLanguage(b.language)))

	if b.style != "" {
		params.Set("style", b.style)
	}

	params.Set("marker_type", strconv.Itoa(int(b.markerType)))

	if b.tenant != "" {
		params.Set("tenant", string(b.tenant))
	}

	return signURL(b.baseURL, fmt.Sprintf("/api/%s/photo/ride", b.version), b.secret, b.expiryDuration, params)
}

func (b *RideRequestBuilder) validate() error {
	if b.baseURL == "" {
		return fmt.Errorf("smapp smappshot: baseURL is required")
	}
	if b.secret == "" {
		return fmt.Errorf("smapp smappshot: signing secret is required")
	}
	if b.expiryDuration <= 0 {
		return fmt.Errorf("smapp smappshot: expiry duration must be positive")
	}
	if b.version == "" {
		return fmt.Errorf("smapp smappshot: version is required")
	}
	if b.language != "" {
		if err := validateLanguage(b.language); err != nil {
			return err
		}
	}
	if b.here == nil {
		if b.origin == nil {
			return fmt.Errorf("smapp smappshot: origin is required when here is not set")
		}
		if len(b.destinations) == 0 {
			return fmt.Errorf("smapp smappshot: at least one destination is required when here is not set")
		}
	}
	if b.markerType != MarkerTypeRideHistory && b.markerType != MarkerTypeLocationShare {
		return fmt.Errorf("smapp smappshot: marker_type must be 0 or 1, got %d", int(b.markerType))
	}
	return nil
}

// PreviewRequestBuilder builds and signs a URL for the /api/{version}/photo/preview endpoint.
//
// Usage:
//
//	url, err := smappshot.NewPreviewRequestBuilder(baseURL, secret, smappshot.V1).
//	    WithCenter(smappshot.Location{Lon: 51.338, Lat: 35.699}).
//	    WithZoom(14).
//	    Build()
type PreviewRequestBuilder struct {
	baseURL        string
	secret         string
	expiryDuration time.Duration
	version        Version
	width          int
	height         int
	language       Language
	style          string
	tenant         Tenant
	center         *Location
	zoom           int
}

// NewPreviewRequestBuilder creates a builder for the preview photo URL.
func NewPreviewRequestBuilder(baseURL string, secret string, version Version) *PreviewRequestBuilder {
	return &PreviewRequestBuilder{
		baseURL:        strings.TrimRight(baseURL, "/"),
		secret:         secret,
		expiryDuration: defaultExpiryDuration,
		version:        version,
	}
}

func (b *PreviewRequestBuilder) WithWidth(w int) *PreviewRequestBuilder {
	b.width = w
	return b
}

func (b *PreviewRequestBuilder) WithHeight(h int) *PreviewRequestBuilder {
	b.height = h
	return b
}

// WithExpiry sets how long the signed URL remains valid. Default is 10 minutes.
func (b *PreviewRequestBuilder) WithExpiry(d time.Duration) *PreviewRequestBuilder {
	b.expiryDuration = d
	return b
}

func (b *PreviewRequestBuilder) WithLanguage(lang Language) *PreviewRequestBuilder {
	b.language = lang
	return b
}

func (b *PreviewRequestBuilder) WithStyle(style string) *PreviewRequestBuilder {
	b.style = style
	return b
}

func (b *PreviewRequestBuilder) WithTenant(tenant Tenant) *PreviewRequestBuilder {
	b.tenant = tenant
	return b
}

func (b *PreviewRequestBuilder) WithCenter(loc Location) *PreviewRequestBuilder {
	b.center = &loc
	return b
}

func (b *PreviewRequestBuilder) WithZoom(zoom int) *PreviewRequestBuilder {
	b.zoom = zoom
	return b
}

// Build validates, signs, and returns the preview photo URL.
func (b *PreviewRequestBuilder) Build() (string, error) {
	if err := b.validate(); err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("width", strconv.Itoa(defaultInt(b.width, 512)))
	params.Set("height", strconv.Itoa(defaultInt(b.height, 285)))
	params.Set("center", formatLocation(*b.center))
	params.Set("zoom", strconv.Itoa(defaultInt(b.zoom, 12)))
	params.Set("language", string(defaultLanguage(b.language)))

	if b.style != "" {
		params.Set("style", b.style)
	}

	if b.tenant != "" {
		params.Set("tenant", string(b.tenant))
	}

	return signURL(b.baseURL, fmt.Sprintf("/api/%s/photo/preview", b.version), b.secret, b.expiryDuration, params)
}

func (b *PreviewRequestBuilder) validate() error {
	if b.baseURL == "" {
		return fmt.Errorf("smapp smappshot: baseURL is required")
	}
	if b.secret == "" {
		return fmt.Errorf("smapp smappshot: signing secret is required")
	}
	if b.expiryDuration <= 0 {
		return fmt.Errorf("smapp smappshot: expiry duration must be positive")
	}
	if b.version == "" {
		return fmt.Errorf("smapp smappshot: version is required")
	}
	if b.language != "" {
		if err := validateLanguage(b.language); err != nil {
			return err
		}
	}
	if b.center == nil {
		return fmt.Errorf("smapp smappshot: center is required")
	}
	return nil
}
