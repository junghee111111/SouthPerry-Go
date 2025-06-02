package enum

import "fmt"

type LoginRecvOp byte

const (
	Pong                  LoginRecvOp = 0x00
	TryLogin              LoginRecvOp = 0x01
	ChannelSelect         LoginRecvOp = 0x04
	RequestCharList       LoginRecvOp = 0x06
	RequestCharNameCheck  LoginRecvOp = 0x07
	RequestCreateChar     LoginRecvOp = 0x08
	LoginScreenTransition LoginRecvOp = 0x42
)

func (o LoginRecvOp) String() string {
	switch o {
	case 0x00:
		return "Pong"
	case 0x01:
		return "TryLogin"
	case 0x04:
		return "ChannelSelect"
	case 0x06:
		return "RequestCharList"
	case 0x07:
		return "RequestCharNameCheck"
	case 0x08:
		return "RequestCreateChar"
	case 0x42:
		return "LoginScreenTransition"
	default:
		return fmt.Sprintf("Unknown(0x%X)", o)
	}
}
