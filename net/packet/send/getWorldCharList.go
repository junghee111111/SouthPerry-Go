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

func BuildGetWorldCharList() []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.WorldCharList.Byte())
	p.WriteByte(0)
	p.WriteUint32(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)

	fmt.Printf("BuildGetWorldCharList Sent : % X\n", p)
	return p.Bytes()
}
