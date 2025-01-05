package reverse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// GetBatch , receives a slice of  Request s and returns Component s of address of location given.
// Does not support type 'frequent' in requests and Does not support type Display option as True
func (c *Client) GetBatch(request BatchReverseRequest) ([]Result, error) {
	return c.GetBatchWithContext(context.Background(), request)
}

// GetBatchWithContext is like GetBatch, but with context.Context support.
// Does not support type 'frequent' in requests and Does not support type Display option as True
func (c *Client) GetBatchWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error) {
	if ctx == nil {
		return nil, fmt.Errorf("smapp reverse geo-code: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-batch-reverse")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("smapp batch reverse geo-code: could not mar request. err: " + err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewBuffer(jsonBody))
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return nil, errors.New("smapp batch reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return nil, fmt.Errorf("smapp batch reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()

	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp batch reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	var results Results
	if response.StatusCode == http.StatusOK {
		err := json.NewDecoder(response.Body).Decode(&results)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return nil, fmt.Errorf("smapp batch reverse geo-code: could not serialize response due to: %s", err.Error())
		}
		responseSpan.End()
		return results.Results, nil
	}

	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return nil, fmt.Errorf("smapp batch reverse geo-code: non 200 status: %d", response.StatusCode)
}

// GetBatchDisplayName , receives a slice of  Request s and returns Component s of address of location given, with only the DisplayName
// Only works when Display is true
func (c *Client) GetBatchDisplayName(request BatchReverseRequest) ([]ResultWithDisplayName, error) {
	return c.GetBatchDisplayNameWithContext(context.Background(), request)
}

// GetBatchDisplayNameWithContext is like GetBatchWithDisplayName, but with context.Context support.
// Only works when Display is true
func (c *Client) GetBatchDisplayNameWithContext(ctx context.Context, request BatchReverseRequest) ([]ResultWithDisplayName, error) {
	if ctx == nil {
		return nil, fmt.Errorf("smapp reverse geo-code: nil context")
	}
	// Start of parent span
	var span trace.Span
	ctx, span = otel.Tracer(c.tracerName).Start(ctx, "get-batch-reverse")
	defer span.End()

	var reqInitSpan trace.Span
	ctx, reqInitSpan = otel.Tracer(c.tracerName).Start(ctx, "request-initialization")

	jsonBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewBuffer(jsonBody))
	if err != nil {
		reqInitSpan.RecordError(err)
		reqInitSpan.End()
		return nil, errors.New("smapp batch reverse geo-code: could not create request. err: " + err.Error())
	}

	params := url.Values{}

	if c.cfg.APIKeySource == config.HeaderSource {
		req.Header.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else if c.cfg.APIKeySource == config.QueryParamSource {
		params.Set(c.cfg.APIKeyName, c.cfg.APIKey)
	} else {
		reqInitSpan.SetStatus(codes.Error, "invalid api key source")
		reqInitSpan.End()
		return nil, fmt.Errorf("smapp batch reverse geo-code: invalid api key source: %s", string(c.cfg.APIKeySource))
	}

	req.Header.Set(version.UserAgentHeader, version.GetUserAgent())

	req.URL.RawQuery = params.Encode()

	reqInitSpan.End()

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("smapp batch reverse geo-code: could not make a request due to this error: %s", err.Error())
	}

	//nolint
	var responseSpan trace.Span
	//nolint
	ctx, responseSpan = otel.Tracer(c.tracerName).Start(ctx, "response-deserialization")

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	var results ResultsWithDisplayName

	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&results)
		if err != nil {
			responseSpan.RecordError(err)
			responseSpan.End()
			return nil, fmt.Errorf("smapp batch reverse geo-code: could not serialize response due to: %s", err.Error())
		}
		responseSpan.End()
		return results.Results, nil
	}

	responseSpan.SetStatus(codes.Error, "non 200 status code")
	responseSpan.End()
	return nil, fmt.Errorf("smapp batch reverse geo-code: non 200 status: %d", response.StatusCode)
}

func (c *Client) GetBatchStructuralResultsWithContext(ctx context.Context, request BatchReverseRequest) ([]StructuralResult, error) {
	results, err := c.GetBatchWithContext(ctx, request)
	if err != nil {
		return nil, err
	}
	structuralResults := make([]StructuralResult, 0)
	for _, result := range results {
		structuralResult := c.convertComponentIntoStructureModel(result.Result.Components)
		structuralResults = append(structuralResults,
			StructuralResult{Result: structuralResult, ID: result.ID})
	}
	return structuralResults, nil
}

func (c *Client) GetBatchStructuralResults(request BatchReverseRequest) ([]StructuralResult, error) {
	return c.GetBatchStructuralResultsWithContext(context.Background(), request)
}
