package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	// ShortenURLHandlerPath is the path for the shorten URL handler
	ShortenURLHandlerPath = "/api/shorten"

	// TargetURL is the URL to shorten
	TargetURL = `{"url":"https://www.example.com"}`
)

func TestShortenURLHandler(t *testing.T) {
	// Create a request body with a URL to shorten
	reqBody := bytes.NewBuffer([]byte(TargetURL))
	req, err := http.NewRequest("POST", ShortenURLHandlerPath, reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShortenURLHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
