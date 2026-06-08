package smappshot

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	testBaseURL = "https://smappshot.example.com"
	testSecret  = "test-secret-key"
)

// ---------------------------------------------------------------------------
// RideRequestBuilder
// ---------------------------------------------------------------------------

func TestRideRequestBuilder_HereMode(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V2).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithLanguage(LanguageEnglish).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path, params, sig := parseSigned(t, rawURL)
	if path != "/api/v2/photo/ride" {
		t.Errorf("path = %q", path)
	}
	if params.Get("here") == "" {
		t.Error("here param missing")
	}
	if params.Get("origin") != "" || params.Get("destinations") != "" {
		t.Error("origin/destinations should be absent in here mode")
	}
	if sig == "" {
		t.Error("sig param missing")
	}
}

func TestRideRequestBuilder_RouteMode(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V2).
		WithOrigin(Location{Lon: 51.338, Lat: 35.699}).
		WithDestinations([]Location{
			{Lon: 51.400, Lat: 35.720},
			{Lon: 51.420, Lat: 35.730},
		}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, params, sig := parseSigned(t, rawURL)
	if params.Get("origin") == "" {
		t.Error("origin param missing")
	}
	if strings.Count(params.Get("destinations"), ";") != 1 {
		t.Errorf("expected 2 destinations separated by ';', got %q", params.Get("destinations"))
	}
	if sig == "" {
		t.Error("sig param missing")
	}
}

func TestRideRequestBuilder_DefaultDimensions(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("width") != "512" {
		t.Errorf("default width = %q, want 512", params.Get("width"))
	}
	if params.Get("height") != "285" {
		t.Errorf("default height = %q, want 285", params.Get("height"))
	}
}

func TestRideRequestBuilder_CustomDimensions(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithWidth(800).
		WithHeight(600).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("width") != "800" {
		t.Errorf("width = %q, want 800", params.Get("width"))
	}
	if params.Get("height") != "600" {
		t.Errorf("height = %q, want 600", params.Get("height"))
	}
}

func TestRideRequestBuilder_DefaultLanguage(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("language") != "fa" {
		t.Errorf("default language = %q, want fa", params.Get("language"))
	}
}

func TestRideRequestBuilder_TenantParam(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithTenant(TenantBalyIQ).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("tenant") != "baly-iq" {
		t.Errorf("tenant = %q, want baly-iq", params.Get("tenant"))
	}
}

func TestRideRequestBuilder_SignatureVerification(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("url.Parse: %v", err)
	}
	q := u.Query()
	sig := q.Get("sig")
	q.Del("sig")

	expected, err := computeSignature(testSecret, u.Path, q)
	if err != nil {
		t.Fatalf("computeSignature error: %v", err)
	}
	if sig != expected {
		t.Errorf("signature mismatch: got %q, want %q", sig, expected)
	}
}

func TestRideRequestBuilder_BaseURLInResult(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(rawURL, testBaseURL) {
		t.Errorf("URL %q does not start with base %q", rawURL, testBaseURL)
	}
}

func TestRideRequestBuilder_CustomExpiry(t *testing.T) {
	rawURL, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithExpiry(30 * time.Minute).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("expires") == "" {
		t.Error("expires param missing")
	}
}

// ---------------------------------------------------------------------------
// RideRequestBuilder validation errors
// ---------------------------------------------------------------------------

func TestRideRequestBuilder_ErrorMissingVersion(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, "").
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err == nil {
		t.Fatal("expected error for missing version")
	}
}

func TestRideRequestBuilder_ErrorMissingSecret(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, "", V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err == nil {
		t.Fatal("expected error for missing secret")
	}
}

func TestRideRequestBuilder_ErrorZeroExpiry(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithExpiry(0).
		Build()
	if err == nil {
		t.Fatal("expected error for zero expiry duration")
	}
}

func TestRideRequestBuilder_ErrorMissingOrigin(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithDestinations([]Location{{Lon: 51.410, Lat: 35.730}}).
		Build()
	if err == nil {
		t.Fatal("expected error for missing origin")
	}
}

func TestRideRequestBuilder_ErrorMissingDestinations(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithOrigin(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err == nil {
		t.Fatal("expected error for missing destinations")
	}
}

func TestRideRequestBuilder_ErrorInvalidLanguage(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithLanguage("zz").
		Build()
	if err == nil {
		t.Fatal("expected error for invalid language")
	}
}

func TestRideRequestBuilder_ErrorInvalidMarkerType(t *testing.T) {
	_, err := NewRideRequestBuilder(testBaseURL, testSecret, V1).
		WithHere(Location{Lon: 51.338, Lat: 35.699}).
		WithMarkerType(99).
		Build()
	if err == nil {
		t.Fatal("expected error for invalid marker_type")
	}
}

// ---------------------------------------------------------------------------
// PreviewRequestBuilder
// ---------------------------------------------------------------------------

func TestPreviewRequestBuilder_Success(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V2).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path, params, sig := parseSigned(t, rawURL)
	if path != "/api/v2/photo/preview" {
		t.Errorf("path = %q", path)
	}
	if params.Get("center") == "" {
		t.Error("center param missing")
	}
	if sig == "" {
		t.Error("sig param missing")
	}
}

func TestPreviewRequestBuilder_DefaultZoom(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("zoom") != "12" {
		t.Errorf("default zoom = %q, want 12", params.Get("zoom"))
	}
}

func TestPreviewRequestBuilder_CustomZoom(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		WithZoom(16).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("zoom") != "16" {
		t.Errorf("zoom = %q, want 16", params.Get("zoom"))
	}
}

func TestPreviewRequestBuilder_TenantParam(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		WithTenant(TenantBalyLBN).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("tenant") != "baly-lbn" {
		t.Errorf("tenant = %q, want baly-lbn", params.Get("tenant"))
	}
}

func TestPreviewRequestBuilder_SignatureVerification(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("url.Parse: %v", err)
	}
	q := u.Query()
	sig := q.Get("sig")
	q.Del("sig")

	expected, err := computeSignature(testSecret, u.Path, q)
	if err != nil {
		t.Fatalf("computeSignature error: %v", err)
	}
	if sig != expected {
		t.Errorf("signature mismatch: got %q, want %q", sig, expected)
	}
}

func TestPreviewRequestBuilder_CustomExpiry(t *testing.T) {
	rawURL, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		WithExpiry(5 * time.Minute).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, params, _ := parseSigned(t, rawURL)
	if params.Get("expires") == "" {
		t.Error("expires param missing")
	}
}

// ---------------------------------------------------------------------------
// PreviewRequestBuilder validation errors
// ---------------------------------------------------------------------------

func TestPreviewRequestBuilder_ErrorMissingVersion(t *testing.T) {
	_, err := NewPreviewRequestBuilder(testBaseURL, testSecret, "").
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		Build()
	if err == nil {
		t.Fatal("expected error for missing version")
	}
}

func TestPreviewRequestBuilder_ErrorMissingCenter(t *testing.T) {
	_, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).Build()
	if err == nil {
		t.Fatal("expected error for missing center")
	}
}

func TestPreviewRequestBuilder_ErrorInvalidLanguage(t *testing.T) {
	_, err := NewPreviewRequestBuilder(testBaseURL, testSecret, V1).
		WithCenter(Location{Lon: 51.338, Lat: 35.699}).
		WithLanguage("xx").
		Build()
	if err == nil {
		t.Fatal("expected error for invalid language")
	}
}

// ---------------------------------------------------------------------------
// formatLocation
// ---------------------------------------------------------------------------

func TestFormatLocation(t *testing.T) {
	got := formatLocation(Location{Lon: 51.338, Lat: 35.699})
	if got != "51.338,35.699" {
		t.Errorf("formatLocation = %q, want 51.338,35.699", got)
	}
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func parseSigned(t *testing.T, rawURL string) (path string, params url.Values, sig string) {
	t.Helper()
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("invalid URL %q: %v", rawURL, err)
	}
	q := u.Query()
	sig = q.Get("sig")
	q.Del("sig")
	return u.Path, q, sig
}
