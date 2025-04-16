package packet

import (
	"awesomeProject/net"
	"fmt"
	"math/rand"
)

func BuildGetHello() []byte {
	// init
	p := &net.MaplePacketWriter{}

	// constants
	ivRecv := []byte{70, 114, 122, byte(rand.Intn(256))}
	ivSend := []byte{82, 48, 120, byte(rand.Intn(256))}
	patchLoc := net.CalcPatchLocation()

	// write packet
	p.WriteShort(uint16(11 + 2 + len(patchLoc)))
	p.WriteShort(291) // KMS static
	p.WriteAsciiString(patchLoc)
	p.Write(ivRecv)
	p.Write(ivSend)
	p.WriteByte(1) // 1 = KMS

	fmt.Printf("GetHello Packet Sent : % X\n", p)
	return p.Bytes()
}
