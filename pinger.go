package main // import "github.com/johnpeterharvey/pinger"

import (
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

func main() {
  log.Print("Starting up...")
  target := os.Getenv("TARGET_URL")
  method := os.Getenv("METHOD")
  interval, err := strconv.Atoi(os.Getenv("INTERVAL"))

  if target == "" || method == "" || err != nil {
    log.Printf("Environment variables were not set, exiting")
    os.Exit(1)
  }

  client := &http.Client{}

  for {
    log.Printf("Trying HTTP %s to %s", method, target)

    req, _ := http.NewRequest(method, target, strings.NewReader("{}"))
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
      log.Printf("Error! Received error while contacting target! %s", err)
    } else {
      log.Printf("Received response %d\n", resp.StatusCode)
      if resp.StatusCode != 200 {
        log.Printf("Error! Received unexpected status code from target! %d - %s", resp.StatusCode, err)
      }
    }

    log.Printf("Sleeping for %d seconds", interval)
    time.Sleep(time.Duration(interval) * time.Second)
  }

  log.Printf("Exiting\n")
}
