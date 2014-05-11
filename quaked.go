package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

type RawPacket struct {
	PacketType uint16
}

type RawControlPacket struct {
	PacketType    uint16
	PacketLength  uint16
	OperationCode uint8
}

type AcceptControlPacket struct {
	PacketType   uint16
	PacketLength uint16
}

func get_raw_packet(buf []byte) *RawPacket {
	packet := &RawPacket{}
	binary.Read(bytes.NewBuffer(buf), binary.BigEndian, packet)
	return packet
}

func get_control_packet(buf []byte) *RawControlPacket {
	packet := &RawControlPacket{}
	binary.Read(bytes.NewBuffer(buf), binary.BigEndian, packet)
	return packet
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:25957")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:26001")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		for {
			buf := make([]byte, 1024)
			_, addr, _ := conn.ReadFromUDP(buf)
			type_packet := get_raw_packet(buf)
			if type_packet.PacketType == 0x8000 {
				control_packet := get_control_packet(buf)
				if control_packet.OperationCode == 0x01 {
					conn.WriteToUDP([]byte{0x80, 0x00, 0x00, 0x09, 0x81, 0x65, 0x00, 0x65, 0x91}, addr)
				}
			}
			fmt.Printf("%s\n", hex.Dump(buf))
		}
	}()

	for {
		buf := make([]byte, 1024)
		_, addr, _ := conn.ReadFromUDP(buf)
		type_packet := get_raw_packet(buf)
		if type_packet.PacketType == 0x8000 {
			control_packet := get_control_packet(buf)
			if control_packet.OperationCode == 0x01 {
				conn.WriteToUDP([]byte{0x80, 0x00, 0x00, 0x09, 0x81, 0x65, 0x65, 0x00, 0x00}, addr)
			}
		}
		fmt.Printf("%s\n", hex.Dump(buf))
	}
}
