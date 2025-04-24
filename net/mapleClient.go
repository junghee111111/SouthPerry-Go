package net

import (
	"SouthPerry/net/encryption"
	"SouthPerry/net/enum"
	"SouthPerry/net/packet/recv"
	"SouthPerry/net/packet/send"
	"bufio"
	"io"
	"log"
	"math/rand"
	"net"
)

type MapleClient struct {
	conn    net.Conn
	KmsRecv *encryption.CryptoManager
	KmsSend *encryption.CryptoManager
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
		KmsRecv: encryption.NewCryptoManager(ivRecv, MapleVersion),
		KmsSend: encryption.NewCryptoManager(ivSend, 0xFFFF-MapleVersion),
	}
}

func HandleClient(c *MapleClient) {
	conn := c.conn
	defer conn.Close()

	acceptClient(c)

	reader := bufio.NewReader(conn)
	for {
		header := make([]byte, 4)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			log.Printf("Failed to read packet length: %v", err)
			return
		}

		packetLength := encryption.DecodePacketLength(c.KmsRecv.Iv[:], c.KmsRecv.VersionIv, header)
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

		decoded := encryption.Decrypt(c.KmsRecv, body)
		opcode := decoded[:1]
		payload := decoded[1:]

		log.Printf("Received packet : %v", packetLength, decoded)
		handlePacket(c, opcode, payload)
	}
}

func acceptClient(c *MapleClient) {
	log.Println("New connection from", c.conn.RemoteAddr())

	patchLoc := CalcPatchLocation()
	helloPacket := send.BuildGetHello(patchLoc, c.ivRecv[:], c.ivSend[:])
	SendRawPacket(c, helloPacket)
}

func SendRawPacket(c *MapleClient, b []byte) {
	_, err := c.conn.Write(b)

	if err != nil {
		log.Printf("Failed to send raw packet: %v", err)
		return
	}
	log.Println("Send raw packet to", c.conn.RemoteAddr())
}

func SendPacket(c *MapleClient, b []byte) {
	_, err := c.conn.Write(encryption.Encrypt(c.KmsSend, b))

	if err != nil {
		log.Printf("Failed to send encrypted packet: %v", err)
		return
	}
	log.Println("Send encrypted packet to", c.conn.RemoteAddr())
}

func handlePacket(c *MapleClient, opcode []byte, payload []byte) {
	_op := enum.LoginRecvOp(opcode[0])
	switch _op {
	case enum.TryLogin:
		log.Println("Opcode 0x01: Client Login Request")
		recv.ParseTryLogin(payload)
		packet := send.BuildGetLoginResult(3)
		SendPacket(c, packet)
	case enum.ChannelSelect:
		log.Println("Opcode 0x04: Channel Select")
	// sendPong(conn)
	case enum.Pong:
		log.Println("Opcode 0x10: Pong")
	default:
		log.Printf("Unhandled opcode: 0x%04X", opcode)
	}
}
