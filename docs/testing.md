# Testing (Mocking)

Each service has a mock client generated with [gomock](https://github.com/golang/mock).

```go
func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearch := search.NewMockSearchClient(ctrl)
	// set expectations on mockSearch...
}
```

For more details see the [gomock documentation](https://github.com/golang/mock).
