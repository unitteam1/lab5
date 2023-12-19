package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	handler := &Server{}
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	expected := `{"message": "hello from production"}`
	actual, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(actual) {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}
