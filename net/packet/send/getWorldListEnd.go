/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package send

import (
	"SouthPerry/net/enum"
	"SouthPerry/net/util"
	"fmt"
)

func BuildGetWorldListEnd() []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.WorldList.Byte())
	p.WriteByte(0xFF)

	fmt.Printf("GetWorldListEnd Packet Sent : % X\n", p)
	return p.Bytes()
}
