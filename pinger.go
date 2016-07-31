package main // import "github.com/johnpeterharvey/pinger"

import (
  "errors"
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

func main() {
  log.Print("Starting up...")

  interval, settings, err := GetSettings()
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }

  client := &http.Client{}

  for {
    DoCall(client, settings)
    log.Printf("Sleeping for %d seconds", interval)
    time.Sleep(time.Duration(interval) * time.Second)
  }
  log.Printf("Exiting\n")
}

func GetSettings() (int, map[string]string, error) {
  target := os.Getenv("TARGET_URL")
  method := os.Getenv("METHOD")
  interval, err := strconv.Atoi(os.Getenv("INTERVAL"))

  if target == "" || method == "" || err != nil || interval < 0 {
    return -1, map[string]string{}, errors.New("Environment variables were not set, returning error")
  }

  return interval, map[string]string{
    "target": target,
    "method": method},
    nil
}

func DoCall(client *http.Client, settings map[string]string) {
  log.Printf("Trying HTTP %s to %s", settings["method"], settings["target"])

  req, _ := http.NewRequest(settings["method"], settings["target"], strings.NewReader("{}"))
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/json")
  resp, err := client.Do(req)

  if err != nil {
    log.Printf("Error! Received error while contacting target! %s", err)
  } else {
    log.Printf("Received response %d\n", resp.StatusCode)
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
      log.Printf("Error! Received unexpected status code from target! %d - %s", resp.StatusCode, err)
    }
    defer resp.Body.Close()
  }
}
