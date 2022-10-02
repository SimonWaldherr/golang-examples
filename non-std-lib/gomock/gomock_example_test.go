package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

// Generates a new server with the a mock client.
func getMockServer(t *testing.T) (*Server, *MockApiClient, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	cli := NewMockApiClient(ctrl)
	server := NewServer(cli)
	return server, cli, ctrl
}

func TestFunctionToTest(t *testing.T) {
	t.Run(
		"Test with mock client - no error",
		func(t *testing.T) {

			// Initialize the mock server and the client.
			// We need to return client seperately because it must be a *MockApiClient and not the ApiClient interface.
			// This is because MockApiClient contains the gomock-specific functions we need to run our tests.
			server, mockCli, _ := getMockServer(t)

			// Sample data
			expectedInput := 2
			expectedOutput := 4

			// Records in mockCli that the provided mock function is called, this is stored in the testing object.
			mockCli.EXPECT().
				// Asserts that ApiDouble is called with the provided input
				ApiDouble(gomock.Eq(expectedInput)).
				// When this is called in the code, these outputs will be returned.
				Return(expectedOutput, nil)

			// Call the function and validate that the output is the expected value and no errors are returned.
			res, err := server.FunctionToTest(expectedInput)

			if res != expectedOutput {
				t.FailNow()
			}

			if err != nil {
				t.FailNow()
			}
		},
	)
}
