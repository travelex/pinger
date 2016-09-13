package main // import "github.com/johnpeterharvey/pinger"

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Print("Starting up...")

	interval, settings, timeToRun, location, err := GetSettings()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	client := &http.Client{}

	if time.Time.IsZero(timeToRun) {
		// Run pinger on a fixed interval
		for {
			log.Printf("Sleeping for %d seconds", interval)
			time.Sleep(time.Duration(interval) * time.Second)
			err := DoCall(client, settings)
			if err != nil {
				log.Print(err)
			}
		}
	} else {
		for {
			// Run pinger at a specific time of day
			duration := GetDurationToWait(timeToRun, location)
			log.Printf("Sleeping for %f seconds", duration.Seconds())
			time.Sleep(duration)
			err = DoCall(client, settings)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

// GetDurationToWait Get duration to sleep until next run
func GetDurationToWait(timeToRun time.Time, timezone *time.Location) time.Duration {
	now := time.Now().In(timezone)
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), timeToRun.Hour(), timeToRun.Minute(), timeToRun.Second(), 0, timezone)

	if !nextRun.After(now) {
		nextRun = time.Date(now.Year(), now.Month(), now.Day()+1, timeToRun.Hour(), timeToRun.Minute(), timeToRun.Second(), 0, timezone)
	}

	return nextRun.Sub(now)
}

// GetSettings Read required settings from environment variables
func GetSettings() (int, map[string]string, time.Time, *time.Location, error) {
	target := os.Getenv("TARGET_URL")
	log.Printf("[Setting] Target: %s", target)
	method := os.Getenv("METHOD")
	log.Printf("[Setting] Method: %s", method)

	interval := 0
	var errInterval error
	if os.Getenv("INTERVAL") != "" {
		interval, errInterval = strconv.Atoi(os.Getenv("INTERVAL"))
		log.Printf("[Setting] Interval: %d", interval)
	}

	timeToRun := time.Time{}
	var errTime error
	if os.Getenv("TIME") != "" {
		timeToRun, errTime = time.Parse("15:04:05", os.Getenv("TIME"))
		log.Printf("[Setting] Time: %s", timeToRun.Format("15:04:05"))
	}

	timeZone := time.UTC
	var errTimezone error
	if os.Getenv("TIMEZONE") != "" {
		timeZone, errTimezone = time.LoadLocation(os.Getenv("TIMEZONE"))
		log.Printf("[Setting] Timezone: %s", timeZone)
	}

	if target == "" || method == "" || errInterval != nil || errTime != nil || errTimezone != nil || interval < 0 {
		return -1, map[string]string{}, time.Time{}, nil, errors.New("Environment variables were not set or could not be parsed, returning error")
	}

	return interval, map[string]string{
			"target": target,
			"method": method},
		timeToRun,
		timeZone,
		nil
}

// DoCall Do an HTTP call out to the server, with the specified properties
func DoCall(client *http.Client, settings map[string]string) error {
	log.Printf("Trying HTTP %s to %s", settings["method"], settings["target"])

	req, err := http.NewRequest(settings["method"], settings["target"], strings.NewReader("{}"))
	if err != nil {
		return fmt.Errorf("Failed to create HTTP request! %s", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return fmt.Errorf("Error! Received error while contacting target! %s", err)
	}
	log.Printf("Received response %d\n", resp.StatusCode)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Error! Received unexpected status code from target! %d", resp.StatusCode)
	}
	return nil
}
