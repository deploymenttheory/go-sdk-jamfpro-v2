package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type DocURLInfo struct {
	URL         string
	ServiceName string
	FilePath    string
	LineNumber  int
}

type TestResult struct {
	URL         string
	StatusCode  int
	Error       error
	ServiceName string
	FilePath    string
	LineNumber  int
}

func main() {
	// Extract all documentation URLs from crud.go files
	docURLs, err := extractDocURLs()
	if err != nil {
		log.Fatalf("Failed to extract documentation URLs: %v", err)
	}

	fmt.Printf("Found %d unique documentation URLs across all services\n", len(docURLs))

	// Test all documentation URLs
	results := testDocURLs(docURLs)

	// Generate report
	if err := generateReport(results); err != nil {
		log.Fatalf("Failed to generate report: %v", err)
	}

	fmt.Printf("\nTesting complete. Results written to tools/doc_url_test_results.log\n")
}

func extractDocURLs() ([]DocURLInfo, error) {
	servicesPath := "../../jamfpro/services"
	urlMap := make(map[string]DocURLInfo)

	// Regular expression to match documentation URLs in comments
	// Matches: // Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
	docURLRegex := regexp.MustCompile(`//.*(?:Jamf Pro API docs|API docs):\s*(https://[^\s]+)`)

	err := filepath.Walk(servicesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileName := filepath.Base(path)
		if !info.IsDir() && fileName == "crud.go" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			serviceName := filepath.Base(filepath.Dir(path))
			scanner := bufio.NewScanner(file)
			lineNumber := 0

			for scanner.Scan() {
				lineNumber++
				line := scanner.Text()

				matches := docURLRegex.FindStringSubmatch(line)
				if len(matches) > 1 {
					url := matches[1]
					// Store unique URLs
					if _, exists := urlMap[url]; !exists {
						urlMap[url] = DocURLInfo{
							URL:         url,
							ServiceName: serviceName,
							FilePath:    path,
							LineNumber:  lineNumber,
						}
					}
				}
			}

			if err := scanner.Err(); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert map to slice and sort by URL
	urls := make([]DocURLInfo, 0, len(urlMap))
	for _, info := range urlMap {
		urls = append(urls, info)
	}

	sort.Slice(urls, func(i, j int) bool {
		return urls[i].URL < urls[j].URL
	})

	return urls, nil
}

func testDocURLs(urls []DocURLInfo) []TestResult {
	results := make([]TestResult, len(urls))
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("\nTesting %d documentation URLs with 20 concurrent workers...\n", len(urls))

	// Create channels for work distribution
	jobs := make(chan struct {
		index   int
		urlInfo DocURLInfo
	}, len(urls))

	// Use a worker pool for concurrent testing
	numWorkers := 20
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				result := TestResult{
					URL:         job.urlInfo.URL,
					ServiceName: job.urlInfo.ServiceName,
					FilePath:    job.urlInfo.FilePath,
					LineNumber:  job.urlInfo.LineNumber,
				}

				// Make HTTP GET request to the documentation URL
				resp, err := client.Get(job.urlInfo.URL)
				if err != nil {
					result.Error = err
					result.StatusCode = 0
				} else {
					result.StatusCode = resp.StatusCode
					resp.Body.Close()
				}

				results[job.index] = result

				// Log progress
				if result.StatusCode != 200 {
					if result.Error != nil {
						fmt.Printf("[%d/%d] ❌ %s - Error: %v\n", job.index+1, len(urls), job.urlInfo.URL, result.Error)
					} else {
						fmt.Printf("[%d/%d] ❌ %s - Status: %d\n", job.index+1, len(urls), job.urlInfo.URL, result.StatusCode)
					}
				} else {
					fmt.Printf("[%d/%d] ✅ %s\n", job.index+1, len(urls), job.urlInfo.URL)
				}
			}
		}()
	}

	// Send jobs to workers
	for i, urlInfo := range urls {
		jobs <- struct {
			index   int
			urlInfo DocURLInfo
		}{i, urlInfo}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	return results
}

func generateReport(results []TestResult) error {
	outputPath := "../../doc_url_test_results.log"
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write header
	fmt.Fprintf(writer, "Jamf Pro Documentation URL Test Results\n")
	fmt.Fprintf(writer, "Generated: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(writer, "%s\n\n", strings.Repeat("=", 80))

	// Count successes and failures
	var successCount, failureCount int
	failures := make([]TestResult, 0)

	for _, result := range results {
		if result.StatusCode == 200 {
			successCount++
		} else {
			failureCount++
			failures = append(failures, result)
		}
	}

	// Write summary
	fmt.Fprintf(writer, "SUMMARY\n")
	fmt.Fprintf(writer, "%s\n", strings.Repeat("-", 80))
	fmt.Fprintf(writer, "Total Documentation URLs Tested: %d\n", len(results))
	fmt.Fprintf(writer, "Successful (200): %d\n", successCount)
	fmt.Fprintf(writer, "Failed (non-200): %d\n\n", failureCount)

	// Write failures
	if len(failures) > 0 {
		fmt.Fprintf(writer, "\nFAILED DOCUMENTATION URLs (Non-200 Responses)\n")
		fmt.Fprintf(writer, "%s\n\n", strings.Repeat("=", 80))

		for i, result := range failures {
			fmt.Fprintf(writer, "%d. URL: %s\n", i+1, result.URL)
			fmt.Fprintf(writer, "   Service: %s\n", result.ServiceName)
			fmt.Fprintf(writer, "   File: %s:%d\n", result.FilePath, result.LineNumber)
			fmt.Fprintf(writer, "   Status Code: %d\n", result.StatusCode)
			if result.Error != nil {
				fmt.Fprintf(writer, "   Error: %v\n", result.Error)
			}
			fmt.Fprintf(writer, "\n")
		}
	}

	// Write successful URLs for reference
	fmt.Fprintf(writer, "\nSUCCESSFUL DOCUMENTATION URLs (200 OK)\n")
	fmt.Fprintf(writer, "%s\n\n", strings.Repeat("=", 80))

	successNum := 1
	for _, result := range results {
		if result.StatusCode == 200 {
			fmt.Fprintf(writer, "%d. %s\n", successNum, result.URL)
			fmt.Fprintf(writer, "   Service: %s\n", result.ServiceName)
			successNum++
		}
	}

	return nil
}
