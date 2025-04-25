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

// BuildGetLoginResult for login failed
func BuildGetLoginResult(result uint32) []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.LoginResult.Byte())
	p.WriteInt(result)
	p.WriteShort(0) // what is this??
	fmt.Printf("send GetLoginResult : % X\n", p)

	return p.Bytes()
}

// BuildGetAuthSuccess for login success
func BuildGetAuthSuccess(result uint32) []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.LoginResult.Byte())
	p.WriteByte(0)
	p.WriteShort(0) // what is this??
	fmt.Printf("send GetLoginResult : % X\n", p)

	return p.Bytes()
}
