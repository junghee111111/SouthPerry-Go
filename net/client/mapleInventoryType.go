/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package client

type MapleInventoryType struct {
	Type byte
	Name string
}

var (
	UNDEFINED = MapleInventoryType{Type: 0, Name: "UNDEFINED"}
	EQUIP     = MapleInventoryType{Type: 1, Name: "EQUIP"}
	USE       = MapleInventoryType{Type: 2, Name: "USE"}
	SETUP     = MapleInventoryType{Type: 3, Name: "SETUP"}
	ETC       = MapleInventoryType{Type: 4, Name: "ETC"}
	CASH      = MapleInventoryType{Type: 5, Name: "CASH"}
	EQUIPPED  = MapleInventoryType{Type: 255, Name: "EQUIPPED"} // -1 is converted to 255 in unsigned representation
)

var AllInventoryTypes = []MapleInventoryType{UNDEFINED, EQUIP, USE, SETUP, ETC, CASH, EQUIPPED}

// GetType returns the byte type of the inventory type.
func (m MapleInventoryType) GetType() byte {
	return m.Type
}

func (m MapleInventoryType) GetBitfieldEncoding() int16 {
	return int16(2) << m.Type
}

func GetByWZName(name string) *MapleInventoryType {
	switch name {
	case "Install":
		return &SETUP
	case "Consume":
		return &USE
	case "Etc":
		return &ETC
	case "Eqp":
		return &EQUIP
	case "Cash", "Pet":
		return &CASH
	default:
		return &UNDEFINED
	}
}
