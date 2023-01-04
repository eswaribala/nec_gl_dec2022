package main

import (
	"fmt"
	"math/rand"
	"necdec2022/day5/models"
	"time"
)

func ProductWrite(ch chan models.Product) {

	for i := 0; i < 10; i++ {
		ch <- models.Product{ProductId: rand.Int31n(10000), Name: "Door", Cost: rand.Int63n(10000),
			DOP: models.Date{Day: 11, Month: 12, Year: 2022}}
		fmt.Println("Successfully Wrote", i, "to channel")

	}
	close(ch)
}

func main() {
	//producer unblocked even it reaches capacity????
	ch := make(chan models.Product, 3)

	go ProductWrite(ch)
	time.Sleep(2 * time.Second)
	for value := range ch {
		fmt.Println("Read Value", value, "from ch")
		time.Sleep(2 * time.Second)
	}
}
