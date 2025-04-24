/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package encryption

func Encrypt(cr *KmsCrypto, b []byte) []byte {
	header := buildPacketHeader(cr.Iv, cr.VersionIv, len(b))

	ivTemp := [4]byte{
		cr.Iv[0], cr.Iv[1], cr.Iv[2], cr.Iv[3],
	}

	UpdateIv(cr)

	for i := 0; i < len(b); i++ {
		input := b[i] & 0xFF
		encrypted := (IVKeys[ivTemp[0]&0xFF] ^ (((0x10*input|(input>>4))>>1)&0x55 | 2*((0x10*input|(input>>4))&0xD5))) & 0xFF

		b[i] = encrypted
		ShuffleIv(input, &ivTemp)
	}

	return append(header[:], b...)
}

func buildPacketHeader(iv [4]byte, version uint16, length int) [4]byte {
	initialIv := uint32(iv[3]) & 0xFF
	initialIv |= (uint32(iv[2]) << 8) & 0xFF00

	initialIv ^= uint32(version)
	processedLength := ((length << 8) & 0xFF00) | (length >> 8)
	ivXor := initialIv ^ uint32(processedLength)

	ret := [4]byte{}
	ret[0] = (byte)((initialIv >> 8) & 0xFF)
	ret[1] = (byte)(initialIv & 0xFF)
	ret[2] = (byte)((ivXor >> 8) & 0xFF)
	ret[3] = (byte)(ivXor & 0xFF)
	return ret
}
