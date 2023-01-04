package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func sleepAccessLink(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("link:%s,status code:%d\n", link, response.StatusCode)

	}

}

func main() {
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
	}

	for _, link := range links {
		time.Sleep(5 * time.Second)
		go sleepAccessLink(link)
	}
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())

}
