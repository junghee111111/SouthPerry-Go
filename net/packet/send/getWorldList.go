/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package send

import (
	"SouthPerry/net/enum"
	"SouthPerry/net/util"
	"SouthPerry/net/world"
	"fmt"
)

var worldKoreanList = []string{
	"스카니아",
	"베라",
	"브로아",
	"카이니",
	"제니스",
	"크로아",
	"아케니아",
	"마르디아",
	"플라나",
	"스티어스",
	"벨로칸",
	"데메토스",
	"옐론드",
	"카스티아",
	"엘니도",
	"윈디아",
	"쥬디스",
	"카디아",
	"갈리시아",
	"칼루나",
}

func BuildGetWorldList(worldId int) []byte {
	p := &util.MaplePacketWriter{}
	tmpWorld := world.NewWorld(fmt.Sprintf("%s", worldKoreanList[worldId]), 1, fmt.Sprintf("이벤트 메세지!"))
	channelSize := 20

	p.WriteByte(enum.WorldList.Byte())

	p.WriteInt(uint32(worldId))

	// 아래부터가 실제로 알아낸 1.2.6 버전의 패킷 구조입니다.
	// 이 버전은 핫/이벤트 플래그, 서버명 페이로드 지원하지 않습니다.
	p.WriteAsciiString(tmpWorld.EventMessage)

	// KMS 1.2.41 Packet structure ##########################
	// p.WriteAsciiString(tmpWorld.Name) // is this event message?
	// p.WriteByte(byte(tmpWorld.Flag))
	// p.WriteAsciiString(tmpWorld.EventMessage)

	// p.WriteShort(100)
	// p.WriteShort(100)
	// KMS 1.2.41 Packet structure ##########################

	p.WriteByte(byte(channelSize)) // channel size

	for i := 1; i <= channelSize; i++ {
		tmpChannelName := fmt.Sprintf("%s-%d", tmpWorld.Name, i)
		if i == 2 {
			tmpChannelName = fmt.Sprintf("%s-20세이상", tmpWorld.Name)
		}

		p.WriteAsciiString(tmpChannelName)
		p.WriteInt(700) // current channel player count (0~1000)
		p.WriteByte(byte(worldId))
		p.WriteShort(uint16(i))
	}

	// 메세지가 있을 때,
	//p.WriteShort(1)
	//p.WriteShort(400) //x-pos
	//p.WriteShort(300)	//y-pos
	//p.WriteAsciiString("ㅋㅋㅋㅋ")

	// 메세지가 없을 때,
	//p.WriteShort(uint16(0))
	//p.WriteInt(uint32(0))

	fmt.Printf("send getServerList : % X \n", p)

	return p.Bytes()
}

func BuildGetWorldListEnd() []byte {
	// init
	p := &util.MaplePacketWriter{}

	// write packet
	p.WriteByte(enum.WorldList.Byte())
	p.WriteByte(0xFF)

	fmt.Printf("GetWorldListEnd Packet Sent : % X\n", p)
	return p.Bytes()
}

/**
func BuildGetWorldList(worldId int) []byte {
	p := &util.MaplePacketWriter{}
	tmpWorld := world.NewWorld(fmt.Sprintf("월드 %d", worldId), 1, fmt.Sprintf("월드 %d Event", worldId))
	channelSize := 19

	p.WriteByte(enum.WorldList.Byte())

	p.WriteByte(byte(worldId))
	p.WriteAsciiString(tmpWorld.Name)
	p.WriteByte(byte(tmpWorld.Flag))
	p.WriteAsciiString(tmpWorld.EventMessage)

	p.WriteByte(0x64)
	p.WriteByte(0)
	p.WriteByte(0x64)
	p.WriteByte(0)

	p.WriteByte(byte(channelSize)) // channel size

	for i := 0; i < channelSize; i++ {
		tmpChannelName := fmt.Sprintf("채널 - %d", i)
		if i == 1 {
			tmpChannelName = "20세이상"
		}
		p.WriteAsciiString(tmpChannelName)
		p.WriteInt(1200) // channel load
		p.WriteByte(byte(worldId))
		p.WriteByte(byte(i)) // channel no.
		p.WriteByte(0)
	}

	fmt.Printf("send getServerList : % X\n", p)

	return p.Bytes()
}
*/
