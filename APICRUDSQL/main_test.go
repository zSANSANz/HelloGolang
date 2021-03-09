package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Home(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	assert.Equal(t, 300, response, "Invalid response code")
}
