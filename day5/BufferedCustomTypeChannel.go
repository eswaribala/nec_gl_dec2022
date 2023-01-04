package main

import (
	"fmt"
	"necdec2022/day5/models"
	"time"
)

func ProductWrite(ch chan models.Product) {

	for i := 0; i < 10; i++ {
		ch <- models.Product{ProductId: 237867, Name: "Door", Cost: 20000,
			DOP: models.Date{Day: 11, Month: 12, Year: 2022}}
		fmt.Println("Successfully Wrote", i, "to channel")
		/*if i == 3 {
			ch = make(chan int, 10)
		}*/
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
