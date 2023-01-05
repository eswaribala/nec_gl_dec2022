package main

import (
	"fmt"
	"io"
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

		handleClock(conn)
	}
}

func handleClock(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05")+"\t")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
