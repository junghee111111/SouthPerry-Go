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
	ivRecv  [4]byte
	ivSend  [4]byte
}

func NewMapleConn(conn net.Conn) *MapleClient {
	ivRecv := [4]byte{70, 114, 122, byte(rand.Intn(256))}
	ivSend := [4]byte{82, 48, 120, byte(rand.Intn(256))}

	log.Printf("IvRecv : %v IvSend : %v", ivRecv, ivSend)
	return &MapleClient{
		conn:    conn,
		ivRecv:  ivRecv,
		ivSend:  ivSend,
		kmsRecv: encryption.NewKmsCrypto(ivRecv, MapleVersion),
		kmsSend: encryption.NewKmsCrypto(ivSend, 0xFFFF-MapleVersion),
	}
}

func HandleClient(c *MapleClient) {
	conn := c.conn
	defer conn.Close()

	accept(c)

	reader := bufio.NewReader(conn)
	for {
		header := make([]byte, 4)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			log.Printf("Failed to read packet length: %v", err)
			return
		}

		packetLength := decodePacketLength(c, header)
		log.Println(" ::: Received packet length", packetLength)

		if packetLength == 0 {
			log.Println(" ::: Packet length is zero")
			return
		}

		body := make([]byte, packetLength)
		_, err = io.ReadFull(reader, body)
		if err != nil {
			log.Printf("Failed to read packet body: %v", err)
			return
		}

		decoded := encryption.Decrypt(c.kmsRecv, body)
		opcode := decoded[:1]
		payload := decoded[1:]

		log.Printf("Received packet : %v", packetLength, decoded)
		handlePacket(opcode, payload)
	}
}

func decodePacketLength(c *MapleClient, stream []byte) uint32 {
	if !isPacketValid(c, stream) {
		log.Printf("Invalid Packet!!  : 0x%04X", stream)
		c.conn.Close()
		return 0
	}
	return getPacketLength(stream)
}

func isPacketValid(c *MapleClient, packetHeader []byte) bool {
	// 여기선 BigEndian 으로 읽어줘야한다..
	rawHeader := binary.BigEndian.Uint32(packetHeader[:4])
	b := make([]byte, 2)
	b[0] = byte((rawHeader >> (8 + 8 + 8)) & 0xFF)
	b[1] = byte((rawHeader >> (8 + 8)) & 0xFF)

	iv := c.kmsRecv.Iv
	version := c.kmsRecv.VersionIv

	return (((b[0] ^ iv[2]) & 0xFF) == byte((version>>8)&0xFF)) &&
		(((b[1] ^ iv[3]) & 0xFF) == byte(version&0xFF))
}

func getPacketLength(packetHeader []byte) uint32 {
	// 여기선 BigEndian 으로 읽어줘야한다..
	rawHeader := binary.BigEndian.Uint32(packetHeader[:4])

	// 00000000 10000000 10000000 00000000
	// 00000000 00000000 00000000 10000000 ^ 00000000 10000000 10000000 00000000
	// 00000000 10000000 10000000 10000000
	pLength := (rawHeader >> 16) ^ (rawHeader & 0xFFFF)

	// 10000000 10000000 00000000 00000000 | 00000000 00000000 10000000 10000000
	pLength = ((pLength << 8) & 0xFF00) | ((pLength >> 8) & 0xFF)
	return pLength
}

func accept(c *MapleClient) {
	log.Println("New connection from", c.conn.RemoteAddr())

	patchLoc := CalcPatchLocation()
	helloPacket := packet.BuildGetHello(patchLoc, c.ivRecv[:], c.ivSend[:])
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

func handlePacket(opcode []byte, payload []byte) {
	_op := opcode[0]
	switch _op {
	case byte(enum.TryLogin):
		log.Println("Opcode 0x01: Client Login Request")
		packet := util.NewPacketReader(payload)
		id := packet.ReadAsciiString()
		password := packet.ReadAsciiString()

		fmt.Printf("    => id : %s | password : %s\n", id, password)

		// handleLogin(payload)
	case byte(enum.ChannelSelect):
		log.Println("Opcode 0x04: Channel Select")
	// sendPong(conn)
	case byte(enum.Pong):
		log.Println("Opcode 0x10: Pong")
	default:
		log.Printf("Unhandled opcode: 0x%04X", opcode)
	}
}
