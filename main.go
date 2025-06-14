package main

import (
	"SouthPerry/db"
	"SouthPerry/db/model"
	mapleNet "SouthPerry/net"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var mongoURI = os.Getenv("MONGO_URI")
var mongoDbName = os.Getenv("MONGO_DB_NAME")
var mariaDBUser = os.Getenv("MARIADB_USER")
var mariaDBPassword = os.Getenv("MARIADB_PASSWORD")
var mariaDBDbName = os.Getenv("MARIADB_DB_NAME")
var mariaDBHost = os.Getenv("MARIADB_HOST")
var mariaDBPort, mariaDBPortConvertErr = strconv.Atoi(os.Getenv("MARIADB_PORT"))

func main() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db.ConnectMariaDB(mariaDBUser, mariaDBPassword, mariaDBHost, mariaDBDbName, mariaDBPort)
	_ = db.MariaDB.AutoMigrate(&model.Account{})
	_ = db.MariaDB.AutoMigrate(&model.Character{})

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
