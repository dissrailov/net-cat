package main

import (
	"log"
	"net"
	"net-cat/internal"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	log.Println("start server on :8080")
	if err != nil {
		log.Fatalf("Error :%s", err)
	}
	go internal.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error :%s", err)
			continue
		}
		go internal.Handle(conn)
	}
}

// TODO: add limit user
