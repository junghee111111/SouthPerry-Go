/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package recv

import (
	"SouthPerry/db/model"
	"SouthPerry/net/util"
)

func ParseRequestCreateChar(payload []byte) model.Character {
	//8 6 0 191 213 193 164 200 241 32 78 0 0 78 117 0 0 130 222 15 0 162 44 16 0 129 91 16 0 240 221 19 0 5 4 8 8
	packet := util.NewPacketReader(payload)

	name := packet.ReadAsciiString()
	face := packet.ReadInt()
	hair := packet.ReadInt()
	top := packet.ReadInt()
	bottom := packet.ReadInt()
	shoes := packet.ReadInt()
	weapon := packet.ReadInt()

	str := uint16(packet.ReadByte())
	dex := uint16(packet.ReadByte())
	innt := uint16(packet.ReadByte())
	luk := uint16(packet.ReadByte())

	var newCharacter = model.Character{}

	newCharacter = model.Character{
		Name:   name,
		Face:   int(face),
		Hair:   int(hair),
		Top:    int(top),
		Bottom: int(bottom),
		Shoes:  int(shoes),
		Weapon: int(weapon),
		Str:    str,
		Dex:    dex,
		Int:    innt,
		Luk:    luk,
		Level:  1,
		Map:    1,
	}

	return newCharacter
}
