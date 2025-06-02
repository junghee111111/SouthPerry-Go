/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package enum

import "fmt"

type LoginSendOp byte

const (
	LoginResult       LoginSendOp = 0x01
	WorldList         LoginSendOp = 0x03
	WorldCharList     LoginSendOp = 0x04
	CheckNameResponse LoginSendOp = 0x06
	AddNewCharacter   LoginSendOp = 0x07

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
	case 0x04:
		return "WorldCharList"
	case 0x06:
		return "CheckNameResponse"
	case 0x07:
		return "AddNewCharacter"
	case 0x0A:
		return "Ping"
	case 0x0F:
		return "GameGuardUpdate"
	default:
		return fmt.Sprintf("Unknown(0x%X)", o)
	}
}

func (o LoginSendOp) Byte() byte {
	return byte(o)
}
