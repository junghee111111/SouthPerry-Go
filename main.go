package main

import (
	"SouthPerry/db"
	mapleNet "SouthPerry/net"
	"context"
	"log"
	"net"
	"os"
	"time"
)

var mongoURI = os.Getenv("MONGO_URI")
var mongoDbName = os.Getenv("MONGO_DB_NAME")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db.ConnectDB(ctx, mongoURI, mongoDbName)

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
