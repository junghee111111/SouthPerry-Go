/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package encryption

import (
	"encoding/binary"
	"log"
)

func Decrypt(cr *KmsCrypto, b []byte) []byte {
	ivTemp := [4]byte{
		cr.Iv[0], cr.Iv[1], cr.Iv[2], cr.Iv[3],
	}

	UpdateIv(cr)

	for i := 0; i < len(b); i++ {
		first := uint32(((b[i] & 0xFF) ^ IVKeys[ivTemp[0]&0xFF]) & 0xFF)
		second := ((first>>1)&0x55 | ((first & 0xD5) << 1)) & 0xFF
		final := ((second << 4) | (second >> 4)) & 0xFF
		b[i] = byte(final)
		ShuffleIv(b[i], &ivTemp)
	}

	return b
}

func DecodePacketLength(iv []byte, version uint16, stream []byte) uint32 {
	if !IsPacketValid(iv, version, stream) {
		log.Printf("Invalid Packet!!  : 0x%04X", stream)
		return 0
	}
	return GetPacketLength(stream)
}

func IsPacketValid(iv []byte, version uint16, packetHeader []byte) bool {
	// 여기선 BigEndian 으로 읽어줘야한다..
	rawHeader := binary.BigEndian.Uint32(packetHeader[:4])
	b := make([]byte, 2)
	b[0] = byte((rawHeader >> (8 + 8 + 8)) & 0xFF)
	b[1] = byte((rawHeader >> (8 + 8)) & 0xFF)

	return (((b[0] ^ iv[2]) & 0xFF) == byte((version>>8)&0xFF)) &&
		(((b[1] ^ iv[3]) & 0xFF) == byte(version&0xFF))
}

func GetPacketLength(packetHeader []byte) uint32 {
	// 여기선 BigEndian 으로 읽어줘야한다..
	rawHeader := binary.BigEndian.Uint32(packetHeader[:4])

	// 00000000 10000000 10000000 00000000
	// 00000000 00000000 00000000 10000000 ^ 00000000 10000000 10000000 00000000
	// 00000000 10000000 10000000 10000000
	pLength := (rawHeader >> 16) ^ (rawHeader & 0xFFFF)

	// 10000000 10000000 00000000 00000000 | 00000000 00000000 10000000 10000000
	pLength = ((pLength << 8) & 0xFF00) | ((pLength >> 8) & 0xFF)
	return pLength
}
