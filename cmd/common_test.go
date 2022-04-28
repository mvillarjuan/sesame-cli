package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSessionIdWithoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/security/login" {
			t.Errorf("Expected to request '/api/v3/security/login', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":"27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f"}`))
	}))
	defer server.Close()

	expectedSessionId := "27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f"
	sessionId, err := getSesionId(server.URL, mock.Anything, mock.Anything)

	assert.Nil(t, err)
	assert.Equal(t, expectedSessionId, sessionId)
}
