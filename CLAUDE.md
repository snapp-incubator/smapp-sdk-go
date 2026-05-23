# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Test with race detector
go test -race $(go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out

# Single package test
go test -race ./services/reverse/... -v

# Lint
golangci-lint run

# Coverage (exclude mocks, matching CI)
grep -v "mock.go" coverage.out > cover.out && go tool cover -func=cover.out

# Dependencies
go mod download && go mod tidy
```

No Makefile — use `go` CLI directly.

## Architecture

Go SDK for Snapp mapping services. Five service clients under `services/`, shared config under `config/`.

### Service Clients

| Package | Purpose |
|---|---|
| `services/reverse` | Lat/lon → address (components, display name, structural) |
| `services/search` | Place search, autocomplete, city lookup |
| `services/area-gateways` | Geographic area polygons and gateways |
| `services/eta` | Estimated time of arrival between points |
| `services/matrix` | Distance/time matrix for source→target pairs |

### Consistent Service Pattern

Every service has:
- `Interface` — defines all operations (enables mocking)
- `Client` struct — concrete implementation
- `NewClient(cfg, version, timeout, ...ClientOption)` — constructor
- `CallOptions` — per-request options (custom headers, OpenTelemetry tracer)
- `*_mock.go` — generated mock via `go.uber.org/mock`
- All operations have `WithContext` variants

### Config

`config.Config` holds: `APIKey`, `Region`, `APIKeySource` (header/query), `APIKeyName`, `APIBaseURL`.

Environment variables: `SMAPP_API_KEY` (required), `SMAPP_API_KEY_SOURCE`, `SMAPP_API_KEY_NAME`, `SMAPP_API_REGION`, `SMAPP_API_BASE_URL`.

Regions: `teh-1` (default), `teh-2`.

### HTTP & Observability

- API key injected via header (`X-Smapp-Key`) or query param (`monshi_key`)
- OpenTelemetry tracing via `otelhttp` — enable with `WithRequestOpenTelemetryTracing(tracerName)` CallOption
- Customizable HTTP transport via `ClientOption`

### Version

`version/version.go` — single constant, used as `User-Agent: smapp-sdk-go/{Version}`.

## Adding a New Service

Follow the pattern of an existing service (e.g., `services/eta`): define `Interface`, implement `Client`, add `options.go`, `call_options.go`, `models.go`, and generate mock with `go.uber.org/mock`.

Regenerate mock after changing an `Interface`:
```bash
mockgen -source services/{name}/{name}.go -destination=services/{name}/{name}_mock.go -mock_names Interface=Mock{Name}Client -package={pkg}
```
Exact command in each `*_mock.go` file header comment.
