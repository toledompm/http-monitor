package trace

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This response doesnt matter!", http.StatusOK)
	}))

	result, err := TraceGetRequest(server.URL)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if result.Response.Status != "200 OK" {
		t.Errorf("Expected: 200 OK, got: %s", result.Response.Status)
	}
	server.Close()
}

func TestErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "This response doesnt matter!", http.StatusInternalServerError)
	}))

	result, err := TraceGetRequest(server.URL)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if result.Response.Status != "500 Internal Server Error" {
		t.Errorf("Expected: 200 OK, got: %s", result.Response.Status)
	}
	server.Close()
}

func TestRequestError(t *testing.T) {
	expectedErrorMessage := "dial tcp [::1]:8080: connect: connection refused"
	result, err := TraceGetRequest("http://localhost:8080/non-existent-url")
	if err == nil {
		t.Errorf("Expected error, got: %s", result.Response.Status)
	}

	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, err.Error())
	}
}
