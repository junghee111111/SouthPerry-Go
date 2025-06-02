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
