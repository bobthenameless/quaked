package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func encodeNulTerminatedString(buf *bytes.Buffer, s string) (n int, err error) {
	n, err = buf.WriteString(s)
	if err != nil {
		return n, err
	}

	err = buf.WriteByte(0)

	return n + 1, err
}

func encodePacketStruct(buf *bytes.Buffer, p interface{}) error {
	ref := reflect.ValueOf(p).Elem()
	for i := 0; i < ref.NumField(); i++ {
		switch t := ref.Field(i).Interface().(type) {
		default:
			return fmt.Errorf("unrecognized type %s in packet definition", t)
		case string:
			_, err := encodeNulTerminatedString(buf, ref.Field(i).String())
			if err != nil {
				return err
			}
		case uint8:
			val := uint8(ref.Field(i).Uint())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case uint16:
			val := uint16(ref.Field(i).Uint())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case uint32:
			val := uint32(ref.Field(i).Uint())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case int8:
			val := int8(ref.Field(i).Int())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case int16:
			val := int16(ref.Field(i).Int())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case int32:
			val := int32(ref.Field(i).Int())
			err := binary.Write(buf, binary.LittleEndian, &val)
			if err != nil {
				return err
			}
		case buint16:
			val := uint16(ref.Field(i).Uint())
			err := binary.Write(buf, binary.BigEndian, &val)
			if err != nil {
				return err
			}
		case buint32:
			val := uint32(ref.Field(i).Uint())
			err := binary.Write(buf, binary.BigEndian, &val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func encodeControlMessage(buf *bytes.Buffer, p interface{}) error {
	payload_buf := &bytes.Buffer{}
	var err error
	var opcode uint8
	switch t := p.(type) {
	default:
		err = fmt.Errorf("unrecognized packet type %s", t)
	case *ConnectionRequestPacket, *ConnectionResponseAcceptPacket, *ConnectionResponseRejectPacket, *ServerInfoRequestPacket, *ServerInfoResponsePacket, *PlayerInfoRequestPacket, *PlayerInfoResponsePacket, *RuleInfoRequestPacket, *RuleInfoResponsePacket:
		err = encodePacketStruct(payload_buf, p)
		switch p.(type) {
		case *ConnectionRequestPacket:
			opcode = ConnectionRequest
		case *ConnectionResponseAcceptPacket:
			opcode = ConnectionResponseAccept
		case *ConnectionResponseRejectPacket:
			opcode = ConnectionResponseReject
		case *ServerInfoRequestPacket:
			opcode = ServerInfoRequest
		case *ServerInfoResponsePacket:
			opcode = ServerInfoResponse
		case *PlayerInfoRequestPacket:
			opcode = PlayerInfoRequest
		case *PlayerInfoResponsePacket:
			opcode = PlayerInfoResponse
		case *RuleInfoRequestPacket:
			opcode = RuleInfoRequest
		case *RuleInfoResponsePacket:
			opcode = RuleInfoResponse
		}
	}

	if err != nil {
		return err
	}

	control_buf := &bytes.Buffer{}
	err = encodePacketStruct(control_buf, &ControlPacketHeader{opcode})

	if err != nil {
		return err
	}

	control_buf_bytes := control_buf.Bytes()
	payload_buf_bytes := payload_buf.Bytes()
	packet_length := len(control_buf_bytes) + len(payload_buf_bytes) + 4

	err = encodePacketStruct(buf, &PacketHeader{ControlMessage, buint16(packet_length)})

	if err != nil {
		return err
	}

	_, err = buf.Write(control_buf_bytes)

	if err != nil {
		return err
	}

	_, err = buf.Write(payload_buf_bytes)

	return err
}

func encodeUnreliableMessage(buf *bytes.Buffer, p interface{}) error {
	return nil
}

func Encode(buf *bytes.Buffer, p interface{}) error {
	switch t := p.(type) {
	default:
		return fmt.Errorf("unrecognized packet type %s", t)
	case *ConnectionRequestPacket, *ConnectionResponseAcceptPacket, *ConnectionResponseRejectPacket, *ServerInfoRequestPacket, *ServerInfoResponsePacket, *PlayerInfoRequestPacket, *PlayerInfoResponsePacket, *RuleInfoRequestPacket, *RuleInfoResponsePacket:
		return encodeControlMessage(buf, p)
	
	case *UnreliableMessagePacket:
		return encodeUnreliableMessage(buf, p)
	}
}
