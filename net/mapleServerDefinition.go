package net

import (
	"math/rand"
	"strconv"
)

var (
	MapleVersion    uint16 = 6 // 1.2.6
	SubVersion      byte   = 1
	RemoveWebCookie byte   = 1

	LoginWaitingLimit       = 5
	LoginProcessingInterval = 500 // (ms)

	versionString *string

	ClientKey  string = randomKey(16)
	DecoderKey string = randomKey(16)
)

func randomKey(n int) string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CalcPatchLocation() string {
	if versionString != nil {
		return *versionString
	}

	// 0000 0000 0000 0000
	// 0111 1111 1111 1111
	ret := int(MapleVersion & 0x7FFF) //  AND연산으로 signed bit 날림
	ret ^= int(RemoveWebCookie) << 15 // 15비트 shift하고 XOR
	ret ^= int(SubVersion) << 16

	tmpVersionString := strconv.Itoa(ret)
	versionString = &tmpVersionString

	return tmpVersionString
}
