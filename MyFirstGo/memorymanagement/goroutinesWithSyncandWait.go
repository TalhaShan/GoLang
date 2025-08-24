package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// websites we want to check
	sites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(sites)) // we will launch one goroutine per site

	for _, site := range sites {
		// launch a goroutine for each site
		go func(s string) {
			defer wg.Done()
			checkWebsite(s)
		}(site)
	}

	// wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Done in %v\n", time.Since(start))
}

func checkWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s -> ERROR: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%s -> %d\n", url, resp.StatusCode)
}

/*
We loop over the list of sites.
For each one, we start a goroutine running checkWebsite.
wg.Add(len(sites)) tells the WaitGroup how many goroutines we’re waiting for.
Each goroutine calls wg.Done() when it finishes.
wg.Wait() blocks until all goroutines are finished.
*/

/*
1. What is sync.WaitGroup?
It’s a type provided by Go’s standard sync package.
A WaitGroup lets you wait for a collection of goroutines to finish.
Think of it like a counter that tracks "how many goroutines are still running".
2. How it works (step by step)
Create one WaitGroup
var wg sync.WaitGroup
→ Now you have a counter, starting at 0.
Add goroutines to wait for
wg.Add(1) // increases counter by 1
Each goroutine you launch should correspond to an Add(1).
If you launch N goroutines, you usually call wg.Add(N) first.
Mark goroutine as done
wg.Done() // decreases counter by 1
Called inside the goroutine once it finishes.
This signals "I’m finished".
Wait until all are done
wg.Wait()
Blocks the main goroutine until the counter goes back to 0.
*/
