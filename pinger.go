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

	interval, settings, timeToRun, loc, err := GetSettings()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	client := &http.Client{}

	if time.Time.IsZero(timeToRun) {
		// run ping on fixed interval only
		for {
			err := DoCall(client, settings)
			if err != nil {
				log.Print(err)
			}
			log.Printf("Sleeping for %d seconds", interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	} else {
		for {
			// run ping at fixed time
			duration := GetDurationToWait(interval, timeToRun, loc)
			log.Printf("Sleeping for %f seconds", duration.Seconds())
			time.Sleep(duration)
			err = DoCall(client, settings)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

// GetDurationToWait Get duration till next run
func GetDurationToWait(interval int, timeToRun time.Time, timezone *time.Location) time.Duration {
	now := time.Now().In(timezone)
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), timeToRun.Hour(), timeToRun.Minute(), timeToRun.Second(), 0, timezone)

	if !nextRun.After(now) {
		nextRun = time.Date(now.Year(), now.Month(), now.Day() + 1, timeToRun.Hour(), timeToRun.Minute(), timeToRun.Second(), 0, timezone)
	}

	return nextRun.Sub(now)
}

// GetSettings Read required settings from environment variables
func GetSettings() (int, map[string]string, time.Time, *time.Location,  error) {
	target := os.Getenv("TARGET_URL")
	method := os.Getenv("METHOD")
	interval, err1 := strconv.Atoi(os.Getenv("INTERVAL"))

	timeToRun := time.Time{}
	var errTime error
	if os.Getenv("TIME") != "" {
		timeToRun, errTime = time.Parse("15:04:05", os.Getenv("TIME"))
	}

	loc := time.UTC
	var errTimezone error
	if os.Getenv("TIMEZONE") != "" {
		loc, errTimezone = time.LoadLocation(os.Getenv("TIMEZONE"))
	}

	if target == "" || method == "" || err1 != nil || errTime != nil || errTimezone != nil || interval < 0 {
		return -1, map[string]string{}, timeToRun, time.UTC,
			errors.New("Environment variables were not set or could not be parsed, returning error")
	}

	return interval, map[string]string{
			"target": target,
			"method": method},
		timeToRun,
		loc,
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
