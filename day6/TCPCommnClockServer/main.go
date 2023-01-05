package main

import (
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
		_, err := io.WriteString(conn, time.Now().Format("05:00:00"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
