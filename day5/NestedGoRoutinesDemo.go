package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//main routine starts
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
	}
	//create the channel
	c := make(chan string)

	for _, link := range links {
		//child routine starts
		//passing channel as parameter
		go accessLinkV2(link, c)
	}
	for l := range c {
		fmt.Println(l)
		go func(link string) {
			time.Sleep(10000)
			go accessLinkV2(link, c)
		}(l)
	}
	//fmt.Println(<-c)
	//main routing executing
	fmt.Println("Completed all routine calls..... ")
}

//receiving channel parameter
func accessLinkV2(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error Occurred", err)
		//writing message in to channel

		c <- link
	} else {
		fmt.Println("Visiting", link)
		//writing message in to channel
		c <- link
	}

}
