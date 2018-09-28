package main

import (
	"fmt"
	"sync"
)

//Fetcher unexported
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	var wg sync.WaitGroup
	var crawl func(string, int, Fetcher, chan string)
	crawl = func(url string, depth int, fetcher Fetcher, ch chan string) {
		wg.Add(1)
		defer wg.Done()
		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			return
		}
		ch <- fmt.Sprintf("found: %s %q\nDepth:%v\n", url, body, depth)

		for _, u := range urls {

			go crawl(u, depth-1, fetcher, ch)

		}
		return
	}
	crawl(url, depth, fetcher, ch)
	wg.Wait()
}

func main() {
	ch := make(chan string)
	go Crawl("https://golang.org/", 5, &fetcher, ch)
	for url := range ch {
		fmt.Println(url)
	}
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	urlList.mx.Lock()
	defer urlList.mx.Unlock()
	_, visited := urlList.list[url]
	if res, ok := f.result[url]; ok && !visited {
		urlList.list[url] = url
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	result map[string]*fakeResult
}

type visitedUrls struct {
	list map[string]string
	mx   sync.Mutex
}

type fakeResult struct {
	body string
	urls []string
}

var urlList = visitedUrls{
	list: make(map[string]string),
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	result: map[string]*fakeResult{
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
		"https://golang.org/cmd/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
				"https://golang.org/pkg/os/linux",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/linux": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/1",
				"https://golang.org/pkg/1",
			},
		},
		"https://golang.org/1": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/2",
				"https://golang.org/pkg/1",
			},
		},
	},
}
