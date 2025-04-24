/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package recv

import (
	"SouthPerry/net/util"
	"fmt"
)

func ParseTryLogin(payload []byte) {
	// init
	packet := util.NewPacketReader(payload)
	id := packet.ReadAsciiString()
	password := packet.ReadAsciiString()

	fmt.Printf("ID: %s\nPWD:%s\n", id, password)
}
