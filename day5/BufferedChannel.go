package main

import (
	"fmt"
	"math/rand"
	"time"
)

func write(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- int(rand.Int31n(10000))
		fmt.Println("Successfully Wrote", i, "to channel")

	}
	close(ch)
}

func main() {
	//producer unblocked even it reaches capacity????
	ch := make(chan int, 3)

	go write(ch)
	time.Sleep(10 * time.Second)

	for value := range ch {
		fmt.Println("Read Value", value, "from ch")
		time.Sleep(2 * time.Second)
	}

}
