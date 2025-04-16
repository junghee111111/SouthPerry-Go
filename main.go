package main

import (
	"SouthPerry/net/packet"
	"bufio"
	"fmt"
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
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from", conn.RemoteAddr())

	helloPacket := packet.BuildGetHello()

	// 2. 클라이언트에게 write
	_, err := conn.Write(helloPacket)
	if err != nil {
		log.Printf("Failed to send hello packet: %v", err)
		return
	}
	log.Println("Sent Hello packet to", conn.RemoteAddr())

	// 3. 클라이언트 응답 받기 루프
	reader := bufio.NewReader(conn)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			log.Printf("Failed to read from client: %v", err)
			return
		}
		fmt.Printf("Raw bytes from client: % X\n", buf[:n])
	}
}
