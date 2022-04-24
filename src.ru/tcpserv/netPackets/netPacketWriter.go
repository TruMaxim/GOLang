package netPackets

import (
	"bytes"
	"encoding/binary"
	"net"
)

type PacketWriter struct {
	Data *bytes.Buffer
	ID   uint16
}

func NewPacketWriter(PackID uint16) PacketWriter {
	packWriter := PacketWriter{new(bytes.Buffer), PackID}
	packWriter.WriteUInt16(PackID)
	return packWriter
}

func (pack *PacketWriter) WriteBool(value bool) error {
	if value {
		return binary.Write(pack.Data, binary.LittleEndian, byte(1))
	}
	return binary.Write(pack.Data, binary.LittleEndian, byte(0))
}

func (pack *PacketWriter) WriteByte(value byte) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteUByte(value uint8) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteInt16(value int16) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteUInt16(value uint16) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteInt32(value int32) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteUInt32(value uint32) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteFloat32(value float32) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteFloat64(value float64) error {
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) WriteString(value string) error {
	chars := []rune(value)
	var strbuf []uint16
	for i := 0; i < len(chars); i++ {
		strbuf = append(strbuf, uint16(chars[i]))
	}
	pack.WriteUInt32(uint32(len(strbuf)))
	return binary.Write(pack.Data, binary.LittleEndian, strbuf)
}

func (pack *PacketWriter) WriteByteArray(value []byte) error {
	pack.WriteUInt32(uint32(len(value)))
	return binary.Write(pack.Data, binary.LittleEndian, value)
}

func (pack *PacketWriter) Send(conn net.Conn) error {
	// Write Length
	lenbuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(lenbuf, uint32(pack.Data.Len()))
	_, err := conn.Write(lenbuf)
	if err != nil {
		return err
	}
	// Write Data
	_, err = conn.Write(pack.Data.Bytes())
	if err != nil {
		return err
	}
	return nil
}
