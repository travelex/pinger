package main_test

import (
  "os"
  "testing"
  "github.com/johnpeterharvey/pinger"
  "github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
  os.Clearenv()

  interval, settings, err := main.GetSettings()
  assert.NotNil(t, err)
  assert.Equal(t, 0, len(settings), "Settings should have failed to fetch")
  assert.Equal(t, -1, interval, "Settings should have failed to fetch")
}
