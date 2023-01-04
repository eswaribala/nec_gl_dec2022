package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func goaccessLink(link string, pChannel chan string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("link:%s,status code:%d\n", link, response.StatusCode)
		pChannel <- link + ":link up and running"
	}

}

func main() {
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
	}
	//create channel
	//unbuffered channel
	channel := make(chan string)
	for _, link := range links {
		//no communication between main and routines
		go goaccessLink(link, channel)
	}
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	for _ = range links {
		fmt.Printf("response:%s\n", <-channel)
	}
	//excess receiver will make main routine keep waiting
	//fmt.Printf("response:%s\n", <-channel)
}
