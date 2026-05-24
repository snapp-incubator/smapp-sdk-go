package smappshot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// RideRequestBuilder builds and signs a URL for the /api/{version}/photo/ride endpoint.
// Either WithHere (single-point mode) or WithOrigin+WithDestinations (route mode) must be called.
//
// Usage:
//
//	url, err := smappshot.NewRideRequestBuilder(baseURL, signingConfig, smappshot.V1).
//	    WithOrigin(smappshot.Location{Lon: 51.338, Lat: 35.699}).
//	    WithDestinations([]smappshot.Location{{Lon: 51.400, Lat: 35.720}}).
//	    Build()
type RideRequestBuilder struct {
	baseURL       string
	signingConfig SigningConfig
	version       Version
	width         int
	height        int
	here          *Location
	origin        *Location
	destinations  []Location
	language      Language
	style         string
	markerType    MarkerType
	tenant        string
}

// NewRideRequestBuilder creates a builder for the ride photo URL.
func NewRideRequestBuilder(baseURL string, signingConfig SigningConfig, version Version) *RideRequestBuilder {
	return &RideRequestBuilder{
		baseURL:       strings.TrimRight(baseURL, "/"),
		signingConfig: signingConfig,
		version:       version,
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

func (b *RideRequestBuilder) WithLanguage(lang Language) *RideRequestBuilder {
	b.language = lang
	return b
}

func (b *RideRequestBuilder) WithStyle(style string) *RideRequestBuilder {
	b.style = style
	return b
}

func (b *RideRequestBuilder) WithMarkerType(mt MarkerType) *RideRequestBuilder {
	b.markerType = mt
	return b
}

func (b *RideRequestBuilder) WithTenant(tenant string) *RideRequestBuilder {
	b.tenant = tenant
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
		params.Set("tenant", b.tenant)
	}

	return signURL(b.baseURL, fmt.Sprintf("/api/%s/photo/ride", b.version), b.signingConfig, params)
}

func (b *RideRequestBuilder) validate() error {
	if err := validateSigningConfig(b.baseURL, b.signingConfig); err != nil {
		return err
	}
	if b.version == "" {
		return fmt.Errorf("smapp smappshot: version is required")
	}
	if b.here == nil {
		if b.origin == nil {
			return fmt.Errorf("smapp smappshot: origin is required when here is not set")
		}
		if len(b.destinations) == 0 {
			return fmt.Errorf("smapp smappshot: at least one destination is required when here is not set")
		}
	}
	if b.language != "" {
		if err := validateLanguage(b.language); err != nil {
			return err
		}
	}
	// MarkerTypeRideHistory=0 is the zero value, so an unset MarkerType is valid by default.
	if b.markerType != MarkerTypeRideHistory && b.markerType != MarkerTypeLocationShare {
		return fmt.Errorf("smapp smappshot: marker_type must be 0 or 1, got %d", int(b.markerType))
	}
	return nil
}

// PreviewRequestBuilder builds and signs a URL for the /api/{version}/photo/preview endpoint.
//
// Usage:
//
//	url, err := smappshot.NewPreviewRequestBuilder(baseURL, signingConfig, smappshot.V1).
//	    WithCenter(smappshot.Location{Lon: 51.338, Lat: 35.699}).
//	    WithZoom(14).
//	    Build()
type PreviewRequestBuilder struct {
	baseURL       string
	signingConfig SigningConfig
	version       Version
	width         int
	height        int
	center        *Location
	zoom          int
	language      Language
	style         string
	tenant        string
}

// NewPreviewRequestBuilder creates a builder for the preview photo URL.
func NewPreviewRequestBuilder(baseURL string, signingConfig SigningConfig, version Version) *PreviewRequestBuilder {
	return &PreviewRequestBuilder{
		baseURL:       strings.TrimRight(baseURL, "/"),
		signingConfig: signingConfig,
		version:       version,
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

func (b *PreviewRequestBuilder) WithCenter(loc Location) *PreviewRequestBuilder {
	b.center = &loc
	return b
}

func (b *PreviewRequestBuilder) WithZoom(zoom int) *PreviewRequestBuilder {
	b.zoom = zoom
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

func (b *PreviewRequestBuilder) WithTenant(tenant string) *PreviewRequestBuilder {
	b.tenant = tenant
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
		params.Set("tenant", b.tenant)
	}

	return signURL(b.baseURL, fmt.Sprintf("/api/%s/photo/preview", b.version), b.signingConfig, params)
}

func (b *PreviewRequestBuilder) validate() error {
	if err := validateSigningConfig(b.baseURL, b.signingConfig); err != nil {
		return err
	}
	if b.version == "" {
		return fmt.Errorf("smapp smappshot: version is required")
	}
	if b.center == nil {
		return fmt.Errorf("smapp smappshot: center is required")
	}
	if b.language != "" {
		if err := validateLanguage(b.language); err != nil {
			return err
		}
	}
	return nil
}

// signURL computes HMAC-SHA256 sig and returns the full signed URL.
// A copy of params is made so the caller's url.Values is not modified.
func signURL(baseURL, path string, cfg SigningConfig, params url.Values) (string, error) {
	// Shallow copy: inner []string slices are shared, but we only Set new keys so this is safe.
	p := make(url.Values, len(params)+1)
	for k, v := range params {
		p[k] = v
	}
	p.Set("expires", strconv.FormatInt(time.Now().UTC().Add(cfg.ExpiryDuration).Unix(), 10))

	sig, err := computeSignature(cfg.Secret, path, p)
	if err != nil {
		return "", fmt.Errorf("smapp smappshot: signing failed: %w", err)
	}

	return baseURL + path + "?" + encodeParamsSorted(p) + "&sig=" + sig, nil
}

// computeSignature returns hex(HMAC-SHA256(secret, "GET\n{path}\n{sorted-params}")).
func computeSignature(secret, path string, params url.Values) (string, error) {
	message := "GET\n" + path + "\n" + encodeParamsSorted(params)
	mac := hmac.New(sha256.New, []byte(secret))
	if _, err := mac.Write([]byte(message)); err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// encodeParamsSorted sorts keys case-insensitively and percent-encodes key=value pairs.
// sig must be excluded before calling.
func encodeParamsSorted(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		for _, v := range params[k] {
			parts = append(parts, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
	}
	return strings.Join(parts, "&")
}

// formatLocation serializes a Location as "{lon},{lat}".
func formatLocation(loc Location) string {
	return strconv.FormatFloat(loc.Lon, 'f', -1, 64) + "," + strconv.FormatFloat(loc.Lat, 'f', -1, 64)
}

func validateLanguage(lang Language) error {
	switch lang {
	case LanguageFarsi, LanguageEnglish, LanguageArabic, LanguageKurdish:
		return nil
	}
	return fmt.Errorf("smapp smappshot: language must be one of fa, en, ar, ku, got %q", string(lang))
}

func validateSigningConfig(baseURL string, cfg SigningConfig) error {
	if baseURL == "" {
		return fmt.Errorf("smapp smappshot: baseURL is required")
	}
	if cfg.Secret == "" {
		return fmt.Errorf("smapp smappshot: signing secret is required")
	}
	if cfg.ExpiryDuration <= 0 {
		return fmt.Errorf("smapp smappshot: expiry duration must be positive")
	}
	return nil
}

// defaultInt returns fallback when val is 0 (zero treated as "not set").
func defaultInt(val, fallback int) int {
	if val == 0 {
		return fallback
	}
	return val
}

func defaultLanguage(lang Language) Language {
	if lang == "" {
		return LanguageFarsi
	}
	return lang
}
