/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package enum

import "fmt"

type LoginSendOp byte

const (
	LoginResult                 LoginSendOp = 0x01
	WorldList                   LoginSendOp = 0x03
	CheckNameResponse           LoginSendOp = 0x06
	Ping                        LoginSendOp = 0x0A
	GameGuardUpdate             LoginSendOp = 0x0F
	PersonIdentifyResult        LoginSendOp = 0x10
	PersonIdentifyPasswordCheck LoginSendOp = 0x11
)

func (o LoginSendOp) String() string {
	switch o {
	case 0x01:
		return "LoginResult"
	case 0x03:
		return "WorldList"
	default:
		return fmt.Sprintf("Unknown Send Op(0x%X)", o)
	}
}

func (o LoginSendOp) Byte() byte {
	return byte(o)
}
