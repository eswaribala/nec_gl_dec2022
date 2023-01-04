package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//Economic Trading
	et := make(chan string)
	//Non Economic Trading
	net := make(chan string)
	//Delta calculation
	dc := make(chan string)

	//anonymous function
	go func(message string, dc chan string) {

		//time.Sleep(1000)
		if strings.ContainsAny(message, "Delta") {
			dc <- "Delta Ready"
		} else {
			dc <- "No Delta Calculations"
		}

	}("ComputeDelta", dc)

	go ecoTrading("Compute ET", et)
	go nonEcoTrading("Compute Leg1 and Leg2", net)

	//select channel receiver
	func() {

		select {
		case msg1 := <-et:
			fmt.Println(msg1)
		case msg2 := <-net:
			fmt.Println(msg2)
		case msg3 := <-dc:
			fmt.Println(msg3)
		}

	}()

	//	var decision string
	//fmt.Println("Do you want to close Trading")
	//fmt.Scanln(&decision)

}

func ecoTrading(message string, et chan string) {

	time.Sleep(5000)
	if len(message) == 0 {
		et <- "Message cannot be empty"
	} else {
		et <- "Eco Trading Done"
	}

}

func nonEcoTrading(message string, net chan string) {

	time.Sleep(7000)
	if strings.HasPrefix(message, "Leg1") {
		net <- "Non Economic Trading Done..."
	} else {
		net <- "Not Done..."
	}

}
