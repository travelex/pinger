package main_test

import (
  "os"
  "testing"
  "github.com/johnpeterharvey/pinger"
  "github.com/stretchr/testify/assert"
)

func TestShouldFailWhenEnvVarsNotSet(t *testing.T) {
  // Given
  os.Clearenv()
  // When
  interval, settings, err := main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")
}

func TestShouldFailWhenNotAllEnvVarsSet(t *testing.T) {
  // Given
  os.Clearenv()
  os.Setenv("TARGET_URL", "test url")
  // When
  interval, settings, err := main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")

  // Given
  os.Clearenv()
  os.Setenv("METHOD", "test method")
  // When
  interval, settings, err = main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")

  // Given
  os.Clearenv()
  os.Setenv("INTERVAL", "10")
  // When
  interval, settings, err = main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")
}

func TestShouldFailWhenIntervalNotParsableInteger(t *testing.T) {
  // Given
  os.Clearenv()
  os.Setenv("INTERVAL", "not an integer!")
  // When
  interval, settings, err := main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")

  // Given
  os.Clearenv()
  os.Setenv("INTERVAL", "-1")
  // When
  interval, settings, err = main.GetSettings()
  // Then
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")
}

func TestShouldReturnSensibleWhenEnvVarsSet(t *testing.T) {
  // Given
  os.Clearenv()
  os.Setenv("TARGET_URL", "test url")
  os.Setenv("METHOD", "method")
  os.Setenv("INTERVAL", "30")
  // When
  interval, settings, err := main.GetSettings()
  // Then
  assert.Nil(t, err)
  assert.Equal(t, 2, len(settings), "Settings should have fetched")
  assert.Equal(t, 30, interval, "Settings should have fetched")
  assert.Equal(t, "method", settings["method"], "Method should be set")
  assert.Equal(t, "test url", settings["target"], "Target should be set")
}
