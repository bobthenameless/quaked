package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

func getNulTerminatedString(reader *bytes.Reader, buf []byte) (n int, err error) {
	for i := 0; i < len(buf); i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return i + 1, err
		}
		buf[i] = b
		if b == byte(0) {
			return i + 1, nil
		}
	}
	return len(buf), errors.New("not enough room in buffer to reach nul")
}

func readPacketStruct(reader *bytes.Reader, p interface{}) error {
	ref := reflect.ValueOf(p).Elem()
	buf := make([]byte, 4096) // used to build up the string fields
	for i := 0; i < ref.NumField(); i++ {
		switch t := ref.Field(i).Interface().(type) {
		default:
			return fmt.Errorf("unrecognized type %s in packet definition", t)
		case string:
			n, err := getNulTerminatedString(reader, buf)
			if err != nil {
				return err
			}
			ref.Field(i).SetString(string(buf[:n-1])) // so long, NUL
		case uint8:
			var val uint8
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetUint(uint64(val))
		case uint16:
			var val uint16
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetUint(uint64(val))
		case uint32:
			var val uint32
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetUint(uint64(val))
		case int8:
			var val int8
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetInt(int64(val))
		case int16:
			var val int16
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetInt(int64(val))
		case int32:
			var val int32
			err := binary.Read(reader, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetInt(int64(val))
		case buint16:
			var val uint16
			err := binary.Read(reader, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetUint(uint64(val))
		case buint32:
			var val uint32
			err := binary.Read(reader, binary.BigEndian, &val)
			if err != nil {
				return err
			}
			ref.Field(i).SetUint(uint64(val))
		}
	}
	return nil
}

func getControlMessage(reader *bytes.Reader, packet_header *PacketHeader) (p interface{}, err error) {
	control_packet_header := &ControlPacketHeader{}
	err = binary.Read(reader, binary.BigEndian, control_packet_header)
	if err != nil {
		return nil, err
	}

	var packet interface{}
	switch control_packet_header.Opcode {
	default:
		return nil, fmt.Errorf("control packet with opcode %d not recognized", control_packet_header.Opcode)
	case ConnectionRequest:
		packet = &ConnectionRequestPacket{}
	case ConnectionResponseAccept:
		packet = &ConnectionResponseAcceptPacket{}
	case ConnectionResponseReject:
		packet = &ConnectionResponseRejectPacket{}
	case ServerInfoRequest:
		packet = &ServerInfoRequestPacket{}
	case ServerInfoResponse:
		packet = &ServerInfoResponsePacket{}
	case PlayerInfoRequest:
		packet = &PlayerInfoRequestPacket{}
	case PlayerInfoResponse:
		packet = &PlayerInfoResponsePacket{}
	case RuleInfoRequest:
		packet = &RuleInfoRequestPacket{}
	case RuleInfoResponse:
		packet = &RuleInfoResponsePacket{}
	}
	err = readPacketStruct(reader, packet)
	if err != nil {
		return nil, err
	}
	return packet, err

}

func getReliableGameMessage(reader *bytes.Reader, packet_header *PacketHeader) (p interface{}, err error) {
	return nil, nil
}

func getReliableGameMessageEnd(reader *bytes.Reader, packet_header *PacketHeader) (p interface{}, err error) {
	return nil, nil
}

func getReliableGameMessageAck(reader *bytes.Reader, packet_header *PacketHeader) (p interface{}, err error) {
	return nil, nil
}

func getUnreliableGameMessage(reader *bytes.Reader, packet_header *PacketHeader) (p interface{}, err error) {
	return nil, nil
}

func Decode(reader *bytes.Reader) (p interface{}, err error) { // TODO: come up with a real type
	packet_header := &PacketHeader{}
	err = binary.Read(reader, binary.BigEndian, packet_header)
	if err != nil {
		return nil, err
	}

	switch packet_header.Type {
	default:
		return nil, fmt.Errorf("packet header code %d not recognized", packet_header.Type)
	case ControlMessage:
		return getControlMessage(reader, packet_header)
	case ReliableGameMessage:
		return getReliableGameMessage(reader, packet_header)
	case ReliableGameMessageEnd:
		return getReliableGameMessageEnd(reader, packet_header)
	case ReliableGameMessageAck:
		return getReliableGameMessageAck(reader, packet_header)
	case UnreliableGameMessage:
		return getReliableGameMessage(reader, packet_header)
	}
}
