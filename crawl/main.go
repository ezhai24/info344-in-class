package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
usage:
	crawl <starting-url>
`

//JobResult stores information about the crawler
type JobResult struct {
	URL   string
	PL    *PageLinks
	Error error
}

func reportResults(result *JobResult, results chan *JobResult) {
	log.Printf("reporting results for %s", result.URL)
	results <- result
}

func startWorking(URLs chan string, results chan *JobResult) {
	for URL := range URLs {
		log.Printf("crawling %s", URL)
		links, err := GetPageLinks(URL)
		result := &JobResult{URL, links, err}

		//crawler hands off results to new go routine
		//can be blocked without deadlocking
		go reportResults(result, results)
	}
}

//numWorkers is the number of worker goroutines
//we will start: begin with just 1 and increase
//to see the benefits of concurrent execution,
//but don't increase beyond the number of concurrent
//socket connections allowed by your OS
const numWorkers = 1

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	//use the first argument as our starting URL
	startingURL := os.Args[1]

	//TODO: build a concurrent web crawler
	//with `numWorkers` worker goroutines,
	//using a channel to pass URLs to fetch
	//form the main goroutine to the workers,
	//and a channel to pass *PageLinks structs
	//from the workers back to the main goroutine.
	//Use the `GetPageLinks()` function in `links.go`
	//from your worker goroutines to fetch links
	//for a given URL.

	//second parameter is how many items can be
	//written into channel before blocking begins
	URLs := make(chan string, 1000)
	results := make(chan *JobResult, 1000)
	seen := map[string]bool{}

	//read from URL channel and start crawling
	for i := 0; i < numWorkers; i++ {
		go startWorking(URLs, results)
	}

	seen[startingURL] = true
	URLs <- startingURL
	// outstandingJobs := 1

	//waits for results to be populated
	for result := range results {
		// outstandin/gJobs--

		//check for error
		if result.Error != nil {
			log.Printf("error crawling %s: %v", result.URL, result.Error)
			continue //breaks out of current for loop iteration, waits for next
		}

		//handle the results
		log.Printf("processing %d links found in %s", len(result.PL.Links), result.URL)
		for _, URL := range result.PL.Links {
			if !seen[URL] {
				seen[URL] = true
				log.Printf("adding %s to the queue", URL)
				URLs <- URL
				// outstandingJobs++
			}
		}

		// if outstandingJobs == 0 {
		// 	log.Println("ALL DONE!")
		// 	return
		// }

		if len(URLs) == 0 {
			log.Println("ALL DONE!")
			return
		}
	}
}
