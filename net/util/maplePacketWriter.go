package util

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

type MaplePacketWriter struct {
	buf bytes.Buffer
}

func (p *MaplePacketWriter) WriteShort(v uint16) {
	binary.Write(&p.buf, binary.LittleEndian, v)
}

func (p *MaplePacketWriter) WriteByte(v byte) {
	p.buf.WriteByte(v)
}

func (p *MaplePacketWriter) WriteInt(v uint32) {
	binary.Write(&p.buf, binary.LittleEndian, v)
}

func (p *MaplePacketWriter) Write(data []byte) {
	p.buf.Write(data)
}

func (p *MaplePacketWriter) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *MaplePacketWriter) WriteAsciiString(s string) {
	encoded := encodeMS949(s)
	p.WriteShort(uint16(len(encoded)))
	p.buf.Write(encoded)
}

func encodeMS949(s string) []byte {
	enc := korean.EUCKR.NewEncoder()
	result, _, _ := transform.String(enc, s)
	return []byte(result)
}
