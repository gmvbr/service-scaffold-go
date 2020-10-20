package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(response, req)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `/`
	if response.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
        response.Body.String(), expected)
	}
}
