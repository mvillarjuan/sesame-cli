package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoCheckInWithoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/employees/ecb42151-1f78-4e23-8f62-40b88df9f2c4/check-in" {
			t.Errorf("Expected request to be sent to /api/v3/employees/ecb42151-1f78-4e23-8f62-40b88df9f2c4/check-in, got %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be application/json, got %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("Authorization") != "Bearer 27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f" {
			t.Errorf("Expected Authorization: Bearer 27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"sessionId":"123456789"}`))
	}))
	sessionId := "27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f"
	userId := "ecb42151-1f78-4e23-8f62-40b88df9f2c4"
	defer server.Close()
	err := doCheckIn(server.URL, sessionId, userId)
	assert.Nil(t, err)
}
