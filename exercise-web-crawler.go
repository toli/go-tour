package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, allUrls map[string]bool, inSignal chan int) {
	if allUrls[url] {
//		fmt.Printf("already saw [%s], returning\n", url)
		inSignal <- 0
		return
	}
//	fmt.Printf("processing [%s]\n", url)
	allUrls[url] = true


	if depth <= 0 {
		inSignal <- 0
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		inSignal <- 0
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	num := len(urls)
	signal := make(chan int, num)
//	fmt.Printf("found [%d] chidren for [%s], going down\n", num, url)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, allUrls, signal)
	}
//	fmt.Printf("will try to wait for [%d] on url [%s]\n", num, url)
	for i:=0;i<num;i++ {
		<-signal
//		fmt.Printf("received [%d] 'done' for url [%s]\n", i, url)
	}
	inSignal <- 0
	return
}

func main() {
	finish := make(chan int, 1)
	allUrls := make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher, allUrls, finish)
	<-finish
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
