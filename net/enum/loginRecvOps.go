package enum

import "fmt"

type LoginRecvOp uint16

const (
	TryLogin      LoginRecvOp = 0x01
	ChannelSelect LoginRecvOp = 0x04
	Pong          LoginRecvOp = 0x0A
)

func (o LoginRecvOp) String() string {
	switch o {
	case 0x01:
		return "TryLogin"
	case 0x04:
		return "ChannelSelect"
	case 0x0A:
		return "Pong"
	default:
		return fmt.Sprintf("Unknown(0x%X)", uint16(o))
	}
}
