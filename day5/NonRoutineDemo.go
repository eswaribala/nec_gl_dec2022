package main

import (
	"fmt"
	"net/http"
)

func accessLink(link string) {
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
	}

}

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
	}

	for _, link := range links {
		accessLink(link)
	}

}
