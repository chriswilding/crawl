package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://www.example.com", "the url to crawl")
	flag.Parse()
	pages := crawl(*urlFlag)
	for _, p := range pages {
		fmt.Printf("%v", p)
	}
}
