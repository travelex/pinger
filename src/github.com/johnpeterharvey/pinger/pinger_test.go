package main_test

import (
  // "net/http/httptest"
  "os"
  "testing"
  "github.com/johnpeterharvey/pinger"
  "github.com/stretchr/testify/assert"
)

// func TestPOST(t *testing.T) {
//   server := httptest.NewServer(
//     http.HandlerFunc(
//       func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "application/json")
//         fmt.Fprintln("")
//       )
//     )
//     defer server.close()
//     url := server.URL
//
//     pinger.main()
// }

func TestSettings(t *testing.T) {
  os.Clearenv()

  interval, settings := main.GetSettings()
  assert.Equal(t, nil, settings, "Settings should have failed to fetch")
  assert.Equal(t, nil, interval, "Settings should have failed to fetch")
}
