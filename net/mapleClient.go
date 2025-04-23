package net

import (
	"SouthPerry/net/enum"
	"SouthPerry/net/packet"
	"SouthPerry/net/util"
	"SouthPerry/net/util/encryption"
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
)

type MapleClient struct {
	conn    net.Conn
	kmsRecv *encryption.KmsCrypto
	kmsSend *encryption.KmsCrypto
	ivRecv  []byte
	ivSend  []byte
}

func NewMapleConn(conn net.Conn) *MapleClient {
	ivRecv := []byte{70, 114, 122, byte(rand.Intn(256))}
	ivSend := []byte{82, 48, 120, byte(rand.Intn(256))}
	return &MapleClient{
		conn:    conn,
		ivRecv:  ivRecv,
		ivSend:  ivSend,
		kmsRecv: encryption.NewKmsCrypto(ivRecv, MapleVersion),
		kmsSend: encryption.NewKmsCrypto(ivSend, (0xFFFF - MapleVersion)),
	}
}

func HandleClient(c *MapleClient) {
	conn := c.conn
	defer conn.Close()

	accept(c)

	reader := bufio.NewReader(conn)
	for {
		header := make([]byte, 2)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			log.Printf("Failed to read packet length: %v", err)
			return
		}
		packetLength := binary.LittleEndian.Uint16(header)
		log.Println(" ::: Received packet length", packetLength)

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

func decodeMapleStream() {

}

func accept(c *MapleClient) {
	log.Println("New connection from", c.conn.RemoteAddr())

	patchLoc := CalcPatchLocation()
	helloPacket := packet.BuildGetHello(patchLoc, c.ivRecv, c.ivSend)
	SendRawPacket(c, helloPacket)
}

func SendRawPacket(c *MapleClient, b []byte) {
	_, err := c.conn.Write(b)

	if err != nil {
		log.Printf("Failed to send hello packet: %v", err)
		return
	}
	log.Println("Sent packet to", c.conn.RemoteAddr())
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
