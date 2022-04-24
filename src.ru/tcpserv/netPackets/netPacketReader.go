package netPackets

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type PacketReader struct {
	Data   *bytes.Buffer
	ID     uint16
	Length uint32
}

func NewPacketReader(buf []byte, packID uint16, length uint32) PacketReader {
	return PacketReader{bytes.NewBuffer(buf), packID, length}
}

func (pack *PacketReader) ReadBool() (bool, error) {
	if pack.Data.Len() < 1 {
		return false, errors.New("EOF end of Packet Data")
	}
	var res byte
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res == 1, nil
}

func (pack *PacketReader) ReadByte() (byte, error) {
	if pack.Data.Len() < 1 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res byte
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadUByte() (uint8, error) {
	if pack.Data.Len() < 1 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res uint8
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadInt16() (int16, error) {
	if pack.Data.Len() < 2 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res int16
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadUInt16() (uint16, error) {
	if pack.Data.Len() < 2 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res uint16
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadInt32() (int32, error) {
	if pack.Data.Len() < 4 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res int32
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadUInt32() (uint32, error) {
	if pack.Data.Len() < 4 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res uint32
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadInt64() (int64, error) {
	if pack.Data.Len() < 8 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res int64
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadUInt64() (uint64, error) {
	if pack.Data.Len() < 8 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res uint64
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadFloat32() (float32, error) {
	if pack.Data.Len() < 4 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res float32
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadFloat64() (float64, error) {
	if pack.Data.Len() < 8 {
		return 0, errors.New("EOF end of Packet Data")
	}
	var res float64
	binary.Read(pack.Data, binary.LittleEndian, &res)
	return res, nil
}

func (pack *PacketReader) ReadByteArray() ([]byte, error) {
	arrLength, err := pack.ReadUInt32()
	if err != nil {
		return nil, err
	}
	if pack.Data.Len() < int(arrLength) {
		return nil, errors.New("EOF end of Packet Data")
	}

	return pack.Data.Next(int(arrLength)), nil
}

func (pack *PacketReader) ReadString() (string, error) {
	if pack.Data.Len() < 4 {
		return "", errors.New("EOF end of Packet Data")
	}

	var strLength int32 = 0
	binary.Read(pack.Data, binary.LittleEndian, &strLength)
	if pack.Data.Len() < int(strLength)*2 {
		return "", errors.New("EOF end of Packet Data")
	}

	res := ""
	var chr int16
	for i := 0; i < int(strLength); i++ {
		binary.Read(pack.Data, binary.LittleEndian, &chr)
		res += string(chr)
	}

	return res, nil
}
