# SmappShot

SmappShot generates **signed static map image URLs** (PNG). No HTTP client is created — the builders produce a URL you pass directly to an `<img>` tag or a `curl` call. Signing is mandatory for Baly deployments.

Base URL for Baly production: `https://staticmap-signed.baly.app`

Import: `github.com/snapp-incubator/smapp-sdk-go/services/smappshot`

## Ride photo

Renders a map with origin/destination markers. Two modes:

- **Route mode** — `WithOrigin` + `WithDestinations` (one or more stops)
- **Here mode** — `WithHere` (single location pin, e.g. driver current position)

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/services/smappshot"
)

func main() {
	const secret = "your-signing-secret"

	// Route mode — Baghdad, Iraq
	rideURL, err := smappshot.NewRideRequestBuilder("https://staticmap-signed.baly.app", secret, smappshot.V2).
		WithOrigin(smappshot.Location{Lat: 33.3152, Lon: 44.3661}).
		WithDestinations([]smappshot.Location{
			{Lat: 33.3600, Lon: 44.4000},
		}).
		WithLanguage(smappshot.LanguageArabic).
		WithTenant(smappshot.TenantBalyIQ).
		WithExpiry(10 * time.Minute).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rideURL)
}
```

## Preview photo

Renders a plain map tile (no markers) centered on a location.

```go
previewURL, err := smappshot.NewPreviewRequestBuilder("https://staticmap-signed.baly.app", secret, smappshot.V2).
	WithCenter(smappshot.Location{Lat: 33.8938, Lon: 35.5018}).
	WithZoom(14).
	WithLanguage(smappshot.LanguageEnglish).
	WithTenant(smappshot.TenantBalyLBN).
	WithExpiry(10 * time.Minute).
	Build()
if err != nil {
	log.Fatal(err)
}
fmt.Println(previewURL)
```

## Builder options

Both `RideRequestBuilder` and `PreviewRequestBuilder` share these options:

| Method | Default | Description |
|---|---|---|
| `WithWidth(int)` | `512` | Image width in pixels |
| `WithHeight(int)` | `285` | Image height in pixels |
| `WithLanguage(Language)` | `LanguageFarsi` | Map label language |
| `WithStyle(string)` | — | Override map style key |
| `WithTenant(Tenant)` | — | Deployment tenant (`TenantBalyIQ`, `TenantBalyLBN`) |
| `WithExpiry(time.Duration)` | `10m` | Signed URL validity window |

Ride-specific:

| Method | Description |
|---|---|
| `WithOrigin(Location)` | Route start point (clears `Here`) |
| `WithDestinations([]Location)` | Route stops (clears `Here`) |
| `WithHere(Location)` | Single-pin mode (clears `Origin`/`Destinations`) |
| `WithMarkerType(MarkerType)` | `MarkerTypeRideHistory` (0) or `MarkerTypeLocationShare` (1) |

Preview-specific:

| Method | Description |
|---|---|
| `WithCenter(Location)` | Map center (**required**) |
| `WithZoom(int)` | Zoom level, default `12` |

## Types

```go
// Versions
smappshot.V1
smappshot.V2

// Languages
smappshot.LanguageFarsi    // "fa"
smappshot.LanguageEnglish  // "en"
smappshot.LanguageArabic   // "ar"
smappshot.LanguageKurdish  // "ku"

// Tenants
smappshot.TenantBalyIQ     // "baly-iq"
smappshot.TenantBalyLBN    // "baly-lbn"

// Marker types
smappshot.MarkerTypeRideHistory    // 0
smappshot.MarkerTypeLocationShare  // 1

// Location — lon comes first on the wire
smappshot.Location{Lat: 33.3152, Lon: 44.3661}
```

## Signing algorithm

URLs are signed with HMAC-SHA256. The message is:

```
GET\n{path}\n{sorted-query-params}
```

Query params are sorted case-insensitively and percent-encoded. The `expires` (Unix timestamp) and `sig` params are appended automatically by `Build()`. The server validates the signature and rejects expired URLs with HTTP 403.
