/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package send

import (
	"SouthPerry/db/model"
	"SouthPerry/net/enum"
	"SouthPerry/net/util"
	"fmt"
)

func BuildResponseCreateChar(c *model.Character) []byte {
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.AddNewCharacter.Byte())

	// if worked, value is 1, if errored, value is 0
	// TODO: determine when the value is 0
	p.WriteByte(1)

	// character info
	p.WriteInt(int(c.ID))
	p.WriteAsciiString(c.Name)
	p.WriteByte(1) // gender
	p.WriteByte(1) // skin color
	p.WriteInt(c.Face)
	p.WriteInt(c.Hair)

	// [0,0,0,0,0,0,0,0]
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)

	p.WriteByte(byte(c.Level))
	p.WriteShort(uint16(c.Job))

	p.WriteShort(10)  // str
	p.WriteShort(10)  // dex
	p.WriteShort(10)  // int
	p.WriteShort(10)  // luk
	p.WriteShort(100) // hp -- SHORT before bigbang
	p.WriteShort(100) // maxhp
	p.WriteShort(100) // mp
	p.WriteShort(100) // maxmp

	p.WriteShort(10) // remaining AP
	p.WriteShort(10) // remaining Sp

	p.WriteInt(0)   // exp
	p.WriteShort(0) // fame
	p.WriteInt(1)   // map Id
	p.WriteByte(0)  // spawn point?

	// character look
	p.WriteByte(1)     // gender
	p.WriteByte(1)     // skin color
	p.WriteInt(c.Face) // face
	p.WriteByte(0)     // mega?
	p.WriteInt(c.Hair)

	// equips
	p.WriteByte(255)

	// masked equips
	p.WriteByte(255)

	// weapon
	p.WriteInt(0)
	p.WriteInt(0)

	// ranking
	if c.Level >= 30 {
		// ranking support
		p.WriteByte(0)

		// TODO: p.writeByte(1)
	} else {
		// ranking not supported
		p.WriteByte(0)
	}

	fmt.Printf("BuildResponseCreateChar Sent : % X\n", p)
	return p.Bytes()
}
