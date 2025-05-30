/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package recv

import "SouthPerry/net/util"

func ParseChannelSelect(payload []byte) (worldId byte, channelId byte) {
	// init
	packet := util.NewPacketReader(payload)

	// drop first 0 byte
	packet.ReadByte()

	iWorldId := packet.ReadByte()
	iChannelId := packet.ReadByte()

	return iWorldId, iChannelId
}
