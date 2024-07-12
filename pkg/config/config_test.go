package config_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/lucasnascimento/ballastlane/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	testDataPath := filepath.Join("testdata", "config_test.json")

	t.Run("ValidConfigFile", func(t *testing.T) {
		cm, err := config.NewConfigManager(testDataPath)
		assert.NoError(t, err)
		assert.Equal(t, testDataPath, cm.Path)
		assert.Equal(t, "tick-value", cm.Config.Tick)
		assert.Equal(t, "tock-value", cm.Config.Tock)
		assert.Equal(t, "bong-value", cm.Config.Bong)
	})

	t.Run("NonexistentConfigFile", func(t *testing.T) {
		_, err := config.NewConfigManager("nonexistent.json")
		assert.Error(t, err)
	})

	t.Run("MalformedConfigFile", func(t *testing.T) {
		malformedTestDataPath := filepath.Join("testdata", "malformed_config_test.json")
		_, err := config.NewConfigManager(malformedTestDataPath)
		assert.Error(t, err)
	})
}

func TestUpdateConfigHandler(t *testing.T) {
	testDataPath := filepath.Join("testdata", "config_test.json")
	cm, err := config.NewConfigManager(testDataPath)
	assert.NoError(t, err)

	// Create a new HTTP request with method POST
	body := bytes.NewBufferString(`{"run_span":0,"tick":"new value","tock":"","bong":""}`)
	req, err := http.NewRequest("POST", "/update-config", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cm.UpdateConfigHandler)

	// Call the handler with our request and recorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"run_span":0,"tick":"new value","tock":"","bong":""}`
	if rr.Body.String() != expected+"\n" { // json.Encoder.Encode adds a newline
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Optionally, check if the ConfigManager's Config field was updated correctly
	if cm.Config.Tick != "new value" {
		t.Errorf("ConfigManager Config was not updated: got %v want %v", cm.Config.Tick, "new value")
	}
}
