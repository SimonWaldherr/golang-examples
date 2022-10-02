package main

type ApiClient interface {
	// When called, the API is expected to double the provided value
	ApiDouble(valueToDouble int) (int, error)
}
