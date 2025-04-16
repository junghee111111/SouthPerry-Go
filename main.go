package main

import (
	"SouthPerry/net/enum"
	"SouthPerry/net/packet"
	"SouthPerry/net/util"
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
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

	_, err := conn.Write(helloPacket)
	if err != nil {
		log.Printf("Failed to send hello packet: %v", err)
		return
	}
	log.Println("Sent Hello packet to", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		header := make([]byte, 2)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			log.Printf("Failed to read packet length: %v", err)
			return
		}
		packetLength := binary.LittleEndian.Uint16(header)

		body := make([]byte, packetLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			log.Printf("Failed to read packet body: %v", err)
			return
		}

		opcode := binary.LittleEndian.Uint16(body[:2])
		payload := body[2:]

		log.Printf("Received packet: Length=%d, Opcode=0x%04X", packetLength, opcode)
		handlePacket(opcode, payload)
	}
}

func handlePacket(opcode uint16, payload []byte) {
	switch opcode {
	case uint16(enum.TryLogin):
		log.Println("Opcode 0x01: Client Login Request")
		packet := util.NewPacketReader(payload)
		id := packet.ReadAsciiString()
		password := packet.ReadAsciiString()

		fmt.Printf("    => id : %s | password : %s\n", id, password)

		// handleLogin(payload)
	case uint16(enum.ChannelSelect):
		log.Println("Opcode 0x04: Channel Select")
	// sendPong(conn)
	case uint16(enum.Pong):
		log.Println("Opcode 0x10: Pong")
	default:

		log.Printf("Unhandled opcode: 0x%04X", opcode)
	}
}
