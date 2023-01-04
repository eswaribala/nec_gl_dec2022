package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func waitAccessLink(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("link:%s,status code:%d\n", link, response.StatusCode)

	}
	wg.Done()
}

func main() {
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
	}
	wg.Add(len(links))
	for _, link := range links {
		//no communication between main and routines
		go waitAccessLink(link)
	}
	wg.Wait()
	fmt.Printf("Number of routines running=%d\n", runtime.NumGoroutine())

}
