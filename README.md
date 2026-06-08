# Smapp Golang SDK

`smapp-sdk-go` is a Go SDK for Smapp services.

## Supported Services

- [x] [Reverse Geocode](docs/reverse.md)
- [x] [Search](docs/search.md)
- [x] [Area Gateways](docs/area-gateways.md)
- [x] [ETA](docs/eta.md)
- [x] [Matrix](docs/matrix.md)
- [x] [SmappShot](docs/smappshot.md) — static map image URL generation
- [ ] Routing

## Installation

Configure your device to access private packages from `github.com/snapp-incubator`, then add to `go.mod`:

```
require (
    github.com/snapp-incubator/smapp-sdk-go v0.9.2
)
```

Download with:

```bash
go mod download
```

> **Note**: Versions before `v0.9.1` are deprecated.

## Configuration

All services (except SmappShot) require a `config` object.

Import: `github.com/snapp-incubator/smapp-sdk-go/config`

### From environment variables

```go
config, err := config.ReadFromEnvironment()
```

| Variable | Default | Description |
|---|---|---|
| `SMAPP_API_KEY` | — | **Required.** API key |
| `SMAPP_API_KEY_SOURCE` | `header` | `header` or `query` |
| `SMAPP_API_KEY_NAME` | `X-Smapp-Key` | Header or query param name |
| `SMAPP_API_REGION` | `teh-1` | `teh-1` or `teh-2` |
| `SMAPP_API_BASE_URL` | `http://smapp-api.apps.inter-dc.okd4.teh-1.snappcloud.io` | API Gateway base URL |

### From code

```go
cfg, err := config.NewDefaultConfig("api-key")
```

### Config options

Both constructors accept options to override defaults:

| Option | Description |
|---|---|
| `WithRegion(string)` | Set region |
| `WithAPIKey(string)` | Set API key |
| `WithAPIBaseURL(string)` | Set custom base URL |
| `WithAPIKeySource(APIKeySource)` | Set key source |
| `WithAPIKeyName(string)` | Set key name |
| `WithPublicURL()` | Use public routes (set region first) |
| `WithInternalURL()` | Use internal routes (set region first) |

```go
cfg, err := config.ReadFromEnvironment(
    config.WithRegion("teh-2"),
    config.WithAPIKey("example-api-key"),
    config.WithPublicURL(),
)
```

## Additional Topics

- [Testing / Mocking](docs/testing.md)
- [OpenTelemetry Tracing](docs/opentelemetry.md)
