package net

import (
	dbEnum "SouthPerry/db/enum"
	"SouthPerry/db/model"
	"SouthPerry/db/service"
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
	conn           net.Conn
	KmsRecv        *encryption.CryptoManager
	KmsSend        *encryption.CryptoManager
	currentWorld   byte
	currentChannel byte
	ivRecv         [4]byte
	ivSend         [4]byte
	account        *model.Account
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
		// log.Println(" ::: Received packet length", packetLength)

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

		log.Printf("Received packet [len:%d] %v", packetLength, decoded)
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
	log.Println("Send raw packet to", c.conn.RemoteAddr(), len(b))
}

func SendPacket(c *MapleClient, b []byte) {
	_, err := c.conn.Write(encryption.Encrypt(c.KmsSend, b))

	if err != nil {
		log.Printf("Failed to send encrypted packet: %v", err)
		return
	}
	log.Println("Send encrypted packet to", c.conn.RemoteAddr(), len(b))
}

func handlePacket(c *MapleClient, opcode []byte, payload []byte) {
	_op := enum.LoginRecvOp(opcode[0])
	// log.Printf("Opcode [%s] %v \n", _op.String(), opcode)
	switch _op {
	case enum.TryLogin:
		email, password := recv.ParseTryLogin(payload)

		/** uncomment if auto account creation is needed */
		service.CreateAccount(email, password)

		respCode, account := service.CheckAccount(email, password)
		if respCode != dbEnum.CheckAccountResp.Success {
			// Login Error
			SendPacket(c, send.BuildGetLoginResult(uint32(respCode)))
		} else {
			// Login Success
			c.account = &account
			SendPacket(c, send.BuildGetAuthSuccess(account))
			for i := 0; i < WorldNum; i++ {
				SendPacket(c, send.BuildGetWorldList(i))
			}
			SendPacket(c, send.BuildGetWorldListEnd())
		}
	case enum.ChannelSelect:
		worldId, channelId := recv.ParseChannelSelect(payload)
		c.currentWorld = worldId
		c.currentChannel = channelId
		SendPacket(c, send.BuildGetWorldCharList())
	case enum.RequestCharNameCheck:
		name := recv.ParseRequestCharNameCheck(payload)
		p := send.BuildResponseCharName(name, service.CheckCharacterName(name))
		SendPacket(c, p)
	case enum.RequestCreateChar:
		newCharacter := recv.ParseRequestCreateChar(payload)
		service.CreateCharacter(c.account.ID, &newCharacter)

		SendPacket(c, send.BuildResponseCreateChar(&newCharacter))
	case enum.Pong:

	case enum.LoginScreenTransition:
		// player login/world/channel screen changed.
		// it works like ping-pong
	default:
		log.Printf("Unhandled opcode: [%s] 0x%02X", _op.String(), opcode)
	}
}
