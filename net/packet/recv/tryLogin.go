/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package recv

import (
	"SouthPerry/net/util"
)

func ParseTryLogin(payload []byte) (email string, password string) {
	// init
	packet := util.NewPacketReader(payload)
	iEmail := packet.ReadAsciiString()
	iPwd := packet.ReadAsciiString()

	return iEmail, iPwd
}
