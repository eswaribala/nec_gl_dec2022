package main

import (
	"flag"
	"fmt"
)

var (
	userName *string
	password *int
)

func init() {
	userName = flag.String("userName", "demo", "default username")
	password = flag.Int("password", 1234, "default password")
}

func main() {
	flag.Parse()
	fmt.Println("UserName", *userName)
	fmt.Println("Password", *password)
}
