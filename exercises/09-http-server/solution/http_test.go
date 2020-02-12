package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPGet200(t *testing.T) {
	expectedResponse := "Success!"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	g := &HTTPGetter{client: &http.Client{}}
	data, err := g.Get(server.URL)
	if err != nil {
		t.Errorf("Unexpected failure state: %s", err)
	}
	actualResponse := data.String()
	
	if expectedResponse != actualResponse {
		t.Errorf("Unexpected response, got: %s, want: %s.", actualResponse, expectedResponse)
	}
}

func TestHTTPGet500(t *testing.T) {
	expectedResponse := "Failure!"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	g := &HTTPGetter{client: &http.Client{}}
	_, err := g.Get(server.URL)
	if err == nil {
		t.Errorf("Should produce failure state")
	}
}

