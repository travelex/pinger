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

  for {
    log.Printf("Trying HTTP %s to %s", method, target)
    if strings.ToUpper(method) == "POST" {
      resp, err := http.Post(target, "application/json", strings.NewReader("{}"))

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
