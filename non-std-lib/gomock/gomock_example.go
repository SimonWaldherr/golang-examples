package main

import "fmt"

// The non-mock implementation of the client
type MyApiClient struct{}

type Server struct {
	Client ApiClient
}

// We can imagine that this function makes some network call during its execution.
func (a *MyApiClient) ApiDouble(valueToDouble int) (int, error) {
	return valueToDouble * 2, nil
}

// Generates a new MyApiClient
func NewMyApiClient() *MyApiClient {
	return &MyApiClient{}
}

// Generates a new server
func NewServer(client ApiClient) *Server {
	return &Server{
		Client: client,
	}
}

func (s *Server) FunctionToTest(val int) (int, error) {
	// this is our imaginary network call
	res, err := s.Client.ApiDouble(val)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func main() {
	cli := NewMyApiClient()
	server := NewServer(cli)
	res, _ := server.FunctionToTest(2)

	fmt.Println(res)
}
