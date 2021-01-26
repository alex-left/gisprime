package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_numHandler(t *testing.T) {
	type cases struct {
		name             string
		path             string
		expectedCode     int
		expectedResponse string
	}
	tests := []cases{
		{name: "test prime", path: "/19", expectedCode: 200, expectedResponse: `{"isPrime":true}`},
		{name: "test not prime", path: "/9", expectedCode: 200, expectedResponse: `{"isPrime":false}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", tt.path, nil)
			responseRecorder := httptest.NewRecorder()
			router := setRouter()
			router.ServeHTTP(responseRecorder, request)
			if responseRecorder.Code != tt.expectedCode {
				t.Errorf("Want status '%d', got '%d'", tt.expectedCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.expectedResponse {
				t.Errorf("Want '%s', got '%s'", tt.expectedResponse, responseRecorder.Body)
			}
		})
	}
}

func Test_healthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status": "ok"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
