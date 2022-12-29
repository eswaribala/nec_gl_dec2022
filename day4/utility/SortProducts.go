package main

import (
	"fmt"
	"math/rand"
	"necdec2022/day4/models"
	"sort"
	"strconv"
)

func main() {

	products := make(models.Products, 10)

	for i := 0; i < len(products); i++ {
		products[i] = models.Product{
			ProductId: rand.Int31n(100000),
			Name:      "Product" + strconv.Itoa(int(rand.Int31n(100000))),
			DOP: models.Date{
				Day:   1 + int8(rand.Int31n(26)),
				Month: 1 + int8(rand.Int31n(11)),
				Year:  1900 + int16(rand.Int31n(122)),
			},
			Cost: int64(rand.Int31n(100000)),
		}
	}

	sort.Sort(products)

	for _, value := range products {
		fmt.Printf("%s\n", value.Name)
	}

}
