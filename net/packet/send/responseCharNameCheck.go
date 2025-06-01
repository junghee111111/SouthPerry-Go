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

func BuildResponseCharName(name string, nameUsed bool) []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.CheckNameResponse.Byte())
	p.WriteAsciiString(name)

	if nameUsed {
		p.WriteByte(1)
	} else {
		p.WriteByte(0)
	}

	fmt.Printf("BuildResponseCharName Sent : % X\n", p)
	return p.Bytes()
}
