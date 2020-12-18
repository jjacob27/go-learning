package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://amazon.com",
		"http://flipkart.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for l := range channel {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, channel)
		}(l)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " not reachable")
	} else {
		fmt.Println(link, "is up!")
	}
	channel <- link
}
