package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func goaccessLink(link string, pChannel chan string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("link:%s,status code:%d\n", link, response.StatusCode)
		pChannel <- "link up and running"
	}

}

func main() {
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
	}

	channel := make(chan string)
	for _, link := range links {
		//no communication between main and routines
		go goaccessLink(link, channel)
	}

	for _, link := range links {
		fmt.Printf("Link:%s,response:%s\n", link, <-channel)
	}
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
}
