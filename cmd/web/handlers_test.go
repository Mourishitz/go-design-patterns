package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestApplication_GetAllCatBreeds(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllCatBreeds)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}
