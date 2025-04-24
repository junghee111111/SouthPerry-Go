/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package enum

import "fmt"

type LoginSendOp byte

const (
	LoginResult LoginSendOp = 0x01
	ServerList  LoginSendOp = 0x03
)

func (o LoginSendOp) String() string {
	switch o {
	case 0x01:
		return "LoginResult"
	case 0x03:
		return "ServerList"
	default:
		return fmt.Sprintf("Unknown Send Op(0x%X)", o)
	}
}

func (o LoginSendOp) Byte() byte {
	return byte(o)
}
