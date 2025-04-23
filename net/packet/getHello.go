package packet

import (
	"SouthPerry/net/util"
	"fmt"
)

func BuildGetHello(patchLoc string, ivRecv []byte, ivSend []byte) []byte {
	// init
	p := &util.MaplePacketWriter{}

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
