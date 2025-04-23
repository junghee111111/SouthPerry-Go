package main

import (
	mapleNet "SouthPerry/net"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8484")
	if err != nil {
		log.Fatalf("Failed to bind due to %v", err)
	}
	log.Println("Listening on port 8484")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to accept new connection due to %v", err)
			continue
		}

		mapleConn := mapleNet.NewMapleConn(conn)
		go mapleNet.HandleClient(mapleConn)
	}
}
