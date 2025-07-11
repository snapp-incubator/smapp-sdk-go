// Code generated by MockGen. DO NOT EDIT.
// Source: services/search/search.go
//
// Generated by this command:
//
//	mockgen -source services/search/search.go -destination=services/search/search_mock.go -mock_names Interface=MockSearchClient -package=search
//

// Package search is a generated GoMock package.
package search

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSearchClient is a mock of Interface interface.
type MockSearchClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchClientMockRecorder
	isgomock struct{}
}

// MockSearchClientMockRecorder is the mock recorder for MockSearchClient.
type MockSearchClientMockRecorder struct {
	mock *MockSearchClient
}

// NewMockSearchClient creates a new mock instance.
func NewMockSearchClient(ctrl *gomock.Controller) *MockSearchClient {
	mock := &MockSearchClient{ctrl: ctrl}
	mock.recorder = &MockSearchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchClient) EXPECT() *MockSearchClientMockRecorder {
	return m.recorder
}

// AutoComplete mocks base method.
func (m *MockSearchClient) AutoComplete(input string, options CallOptions) ([]Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoComplete", input, options)
	ret0, _ := ret[0].([]Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AutoComplete indicates an expected call of AutoComplete.
func (mr *MockSearchClientMockRecorder) AutoComplete(input, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoComplete", reflect.TypeOf((*MockSearchClient)(nil).AutoComplete), input, options)
}

// AutoCompleteWithContext mocks base method.
func (m *MockSearchClient) AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoCompleteWithContext", ctx, input, options)
	ret0, _ := ret[0].([]Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AutoCompleteWithContext indicates an expected call of AutoCompleteWithContext.
func (mr *MockSearchClientMockRecorder) AutoCompleteWithContext(ctx, input, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoCompleteWithContext", reflect.TypeOf((*MockSearchClient)(nil).AutoCompleteWithContext), ctx, input, options)
}

// Details mocks base method.
func (m *MockSearchClient) Details(placeId string, options CallOptions) (Detail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Details", placeId, options)
	ret0, _ := ret[0].(Detail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Details indicates an expected call of Details.
func (mr *MockSearchClientMockRecorder) Details(placeId, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Details", reflect.TypeOf((*MockSearchClient)(nil).Details), placeId, options)
}

// DetailsWithContext mocks base method.
func (m *MockSearchClient) DetailsWithContext(ctx context.Context, placeId string, options CallOptions) (Detail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetailsWithContext", ctx, placeId, options)
	ret0, _ := ret[0].(Detail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetailsWithContext indicates an expected call of DetailsWithContext.
func (mr *MockSearchClientMockRecorder) DetailsWithContext(ctx, placeId, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetailsWithContext", reflect.TypeOf((*MockSearchClient)(nil).DetailsWithContext), ctx, placeId, options)
}

// GetCities mocks base method.
func (m *MockSearchClient) GetCities(options CallOptions) ([]City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCities", options)
	ret0, _ := ret[0].([]City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCities indicates an expected call of GetCities.
func (mr *MockSearchClientMockRecorder) GetCities(options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCities", reflect.TypeOf((*MockSearchClient)(nil).GetCities), options)
}

// GetCitiesWithContext mocks base method.
func (m *MockSearchClient) GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCitiesWithContext", ctx, options)
	ret0, _ := ret[0].([]City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCitiesWithContext indicates an expected call of GetCitiesWithContext.
func (mr *MockSearchClientMockRecorder) GetCitiesWithContext(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCitiesWithContext", reflect.TypeOf((*MockSearchClient)(nil).GetCitiesWithContext), ctx, options)
}

// SearchCity mocks base method.
func (m *MockSearchClient) SearchCity(input string, options CallOptions) ([]City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCity", input, options)
	ret0, _ := ret[0].([]City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCity indicates an expected call of SearchCity.
func (mr *MockSearchClientMockRecorder) SearchCity(input, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCity", reflect.TypeOf((*MockSearchClient)(nil).SearchCity), input, options)
}

// SearchCityWithContext mocks base method.
func (m *MockSearchClient) SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCityWithContext", ctx, input, options)
	ret0, _ := ret[0].([]City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCityWithContext indicates an expected call of SearchCityWithContext.
func (mr *MockSearchClientMockRecorder) SearchCityWithContext(ctx, input, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCityWithContext", reflect.TypeOf((*MockSearchClient)(nil).SearchCityWithContext), ctx, input, options)
}
