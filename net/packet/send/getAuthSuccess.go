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
	"time"
)

// BuildGetAuthSuccess for login success
func BuildGetAuthSuccess(acc model.Account) []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.LoginResult.Byte())
	p.WriteByte(0)

	// account info
	p.WriteUint32(uint32(acc.AccId))
	if acc.Sex == false {
		p.WriteByte(0)
	} else {
		p.WriteByte(1) // female
	}
	p.WriteByte(0)                // is gm
	p.WriteAsciiString(acc.Email) // email

	p.WriteUint32(3) // 뭔지 모름
	p.WriteByte(0)
	p.WriteByte(0)
	p.WriteByte(0)                      // chat banned
	p.WriteLong(time.Now().UnixMilli()) // chat banned time
	p.WriteAsciiString("")              // ?
	p.WriteAsciiString("")              // ?

	fmt.Printf("send GetAuthSuccess : % X\n", p)

	return p.Bytes()
}
