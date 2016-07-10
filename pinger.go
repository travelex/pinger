package main

import (
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

func main() {
  log.Print("Starting up...\n")
  target := os.Getenv("TARGET_URL")
  method := os.Getenv("METHOD")
  interval, _ := strconv.Atoi(os.Getenv("INTERVAL"))

  if method == "POST" {
    resp, _ := http.Post(target, "application/json", strings.NewReader("{}"))

    log.Printf("Received response %d\n", resp.StatusCode)
  }

  time.Sleep(time.Duration(interval) * time.Second)
  log.Printf("Exiting\n")
}
