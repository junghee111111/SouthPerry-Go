package util

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io"
)

type MaplePacketReader struct {
	buf *bytes.Buffer
}

func NewPacketReader(data []byte) *MaplePacketReader {
	return &MaplePacketReader{buf: bytes.NewBuffer(data)}
}

func (r *MaplePacketReader) ReadByte() byte {
	b, _ := r.buf.ReadByte()
	return b
}

func (r *MaplePacketReader) ReadShort() uint16 {
	var val uint16
	binary.Read(r.buf, binary.LittleEndian, &val)
	return val
}

func (r *MaplePacketReader) ReadInt() uint32 {
	var val uint32
	binary.Read(r.buf, binary.LittleEndian, &val)
	return val
}

func (r *MaplePacketReader) ReadAsciiString() string {
	len := r.ReadShort()
	b := make([]byte, len)
	decoder := korean.EUCKR.NewDecoder()

	binary.Read(r.buf, binary.LittleEndian, &b)
	result, _ := io.ReadAll(transform.NewReader(bytes.NewReader(b), decoder))

	return string(result)
}
