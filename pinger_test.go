package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldFailWhenEnvVarsNotSet(t *testing.T) {
	// Given
	os.Clearenv()
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")
}

func TestShouldFailWhenNotAllEnvVarsSet(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("TARGET_URL", "test url")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")

	// Given
	os.Clearenv()
	os.Setenv("METHOD", "test method")
	// When
	interval, settings, timeInput, timezone, err = GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")

	// Given
	os.Clearenv()
	os.Setenv("INTERVAL", "10")
	// When
	interval, settings, timeInput, timezone, err = GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")
}

func TestShouldFailWhenIntervalNotParsableInteger(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("INTERVAL", "not an integer!")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")

	// Given
	os.Clearenv()
	os.Setenv("INTERVAL", "-1")
	// When
	interval, settings, timeInput, timezone, err = GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")
}

func TestShouldReturnSensibleWhenEnvVarsSet(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("TARGET_URL", "test url")
	os.Setenv("METHOD", "method")
	os.Setenv("INTERVAL", "30")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.Nil(t, err)
	assert.Equal(t, 2, len(settings), "Settings should have fetched")
	assert.Equal(t, 30, interval, "Settings should have fetched")
	assert.Equal(t, "method", settings["method"], "Method should be set")
	assert.Equal(t, "test url", settings["target"], "Target should be set")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, time.UTC, timezone, "Timezone should have been default value")
}

func TestShouldReturnSensibleWhenTimeAndTimezoneEnvVarsSet(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("TARGET_URL", "test url")
	os.Setenv("METHOD", "method")
	os.Setenv("TIME", "11:00:00")
	os.Setenv("TIMEZONE", "Europe/London")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	expectTime, err1 := time.Parse("15:04:05", "11:00:00")
	expectLocation, err2 := time.LoadLocation("Europe/London")
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, 2, len(settings), "Settings should have fetched")
	assert.Equal(t, 0, interval, "Settings should have fetched")
	assert.Equal(t, "method", settings["method"], "Method should be set")
	assert.Equal(t, "test url", settings["target"], "Target should be set")
	assert.Equal(t, expectTime, timeInput, "Time should have been set")
	assert.Equal(t, expectLocation, timezone, "Timezone should have been set")
}

func TestShouldFailWhenUnparsableTimeSet(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("TARGET_URL", "test url")
	os.Setenv("METHOD", "method")
	os.Setenv("TIME", "unparsable")
	os.Setenv("TIMEZONE", "Europe/London")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")
}

func TestShouldFailWhenUnparsableTimezoneSet(t *testing.T) {
	// Given
	os.Clearenv()
	os.Setenv("TARGET_URL", "test url")
	os.Setenv("METHOD", "method")
	os.Setenv("TIME", "11:00:00")
	os.Setenv("TIMEZONE", "invalid")
	// When
	interval, settings, timeInput, timezone, err := GetSettings()
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
	assert.Equal(t, -1, interval, "Settings should have failed to fetch")
	assert.Equal(t, time.Time{}, timeInput, "Time should have been default value")
	assert.Equal(t, (*time.Location)(nil), timezone, "Timezone should have been default value")
}

func createServer(code int) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
	}))

	httpClient := &http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}

	return server, httpClient
}

func TestShouldHandleHTTPFailure(t *testing.T) {
	// Given
	settings := map[string]string{
		"target": "http://example.com",
		"method": "POST"}
	server, client := createServer(500)
	defer server.Close()

	// When
	err := DoCall(client, settings)

	// Then
	assert.NotNil(t, err)
}

func TestShouldHandleHTTPSuccess(t *testing.T) {
	// Given
	settings := map[string]string{
		"target": "http://example.com",
		"method": "POST"}
	server, client := createServer(200)
	defer server.Close()

	// When
	err := DoCall(client, settings)

	// Then
	assert.Nil(t, err)
}
