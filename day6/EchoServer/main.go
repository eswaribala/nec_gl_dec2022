package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3333")

	if err != nil {
		log.Println("Connection Error")
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("Connection Established", conn.LocalAddr().String())
		if err != nil {
			log.Fatal(err)
			continue
		}

		handleClientInput(conn)
	}
}

func handleClientInput(conn net.Conn) {
	defer conn.Close()
	for {
		input := bufio.NewScanner(conn)
		if input.Scan() {
			echo(conn, input.Text(), 500)
		}
		
		time.Sleep(1 * time.Second)
	}
}

func echo(conn net.Conn, input string, delay time.Duration) {

	fmt.Fprintf(conn, "The Message received from the client%s", input)
	time.Sleep(delay * time.Second)
}
