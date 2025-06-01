/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package recv

import "SouthPerry/net/util"

func ParseRequestCharNameCheck(payload []byte) string {
	// init
	packet := util.NewPacketReader(payload)

	return packet.ReadAsciiString()
}
