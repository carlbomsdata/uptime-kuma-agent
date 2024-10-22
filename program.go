package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const logFile = "log.txt"

// Setup logging
func setupLogging() *log.Logger {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}

	logger := log.New(file, "", log.LstdFlags)
	return logger
}

// Perform ping test and extract the average ping time
func ping(isp string) string {
	cmd := exec.Command("ping", "-c", "1", isp)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "N/A"
	}

	// Look for "avg" in ping result
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "avg") {
			parts := strings.Split(line, "/")
			if len(parts) >= 5 {
				return parts[4] // avg ping value
			}
		}
	}
	return "N/A"
}

func main() {
	// Set up logging
	logger := setupLogging()

	// Parse command-line arguments
	isp := flag.String("isp", os.Getenv("ISP"), "ISP server to ping")
	baseURL := flag.String("base_url", os.Getenv("BASE_URL"), "Base URL for the HTTP request")
	interval := flag.Int("interval", 0, "Interval time in seconds to rerun the script")
	flag.Parse()

	// Log environment variables
	logger.Println("START")
	logger.Printf("ISP: %s\n", *isp)
	logger.Printf("BASE_URL: %s\n", *baseURL)

	// Validate ISP and BASE_URL
	if *isp == "" {
		logger.Println("No ISP provided. Please provide an ISP to ping.")
		return
	}
	if *baseURL == "" {
		logger.Println("No BASE_URL provided. Please provide a BASE_URL for the HTTP request.")
		return
	}

	// Infinite loop for running the test at intervals
	for {
		// Perform the ping test
		pingResult := ping(*isp)

		// Construct the request URL with the dynamic ping value
		url := fmt.Sprintf("%s?status=up&msg=OK&ping=%s", *baseURL, pingResult)
		logger.Printf("FULL_URL: %s\n", url)

		// Execute the HTTP request
		resp, err := http.Get(url)
		if err != nil {
			logger.Printf("Failed to execute HTTP request: %v\n", err)
		} else {
			defer resp.Body.Close()
			logger.Printf("Response status: %d\n", resp.StatusCode)
		}

		// If no interval is provided, run only once
		if *interval == 0 {
			break
		} else {
			logger.Printf("Sleeping for %d seconds\n", *interval)
			time.Sleep(time.Duration(*interval) * time.Second)
		}
	}
}
