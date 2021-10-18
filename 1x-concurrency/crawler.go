package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(wg *sync.WaitGroup, url string, depth int, fetcher Fetcher) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	defer wg.Done()

	attemptedURLs.Mu.Lock()
	defer attemptedURLs.Mu.Unlock()

	for _, u := range attemptedURLs.Links {
		if url == u {
			fmt.Println("this url was already processed: ", url)
			return
		}
	}

	// processing a fresh url
	fmt.Println("processing: ", url)

	attemptedURLs.Links = append(attemptedURLs.Links, url)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// u := u

		fmt.Printf("fetching %s (parallel) \n", u)

		wg.Add(1)
		go Crawl(wg, u, depth-1, fetcher)
	}
	return
}

type FetchStatus struct {
	Links []string
	Mu    sync.Mutex
}

var attemptedURLs FetchStatus

func main() {
	attemptedURLs = FetchStatus{}

	var wg sync.WaitGroup

	// fmt.Printf("%v", attemptedURLs.Links)
	wg.Add(1)
	go Crawl(&wg, "https://golang.org/", 4, fetcher)
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

	fmt.Println(attemptedURLs.Links)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
