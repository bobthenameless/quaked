package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/the-barmstrong/quaked/packet"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":26000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on %s", conn.LocalAddr().String())
	defer conn.Close()
	for {
		buf := make([]byte, 4096)
		_, remote, _ := conn.ReadFromUDP(buf)
		recv_buf := bytes.NewReader(buf)
		recv_packet, err := packet.Decode(recv_buf)
		if err != nil {
			log.Fatal(err)
		}
		switch recv_packet.(type) {
		default:
			continue
		case *packet.ConnectionRequestPacket:
			fmt.Println("Saw connection attempt")
		case *packet.ServerInfoRequestPacket:
			send_buf := &bytes.Buffer{}
			send_packet := &packet.ServerInfoResponsePacket{conn.LocalAddr().String(), "HOSTNAME", "e1m1", 0, 4, 3}
			err = packet.Encode(send_buf, send_packet)
			if err != nil {
				log.Fatal(err)
			}
			send_bytes := send_buf.Bytes()
			fmt.Println("%d", send_bytes)
			sent_len, err := conn.WriteToUDP(send_bytes, remote)
			if err != nil {
				log.Fatal(err)
			}
			if sent_len != len(send_bytes) {
				log.Fatal(errors.New("WriteToUDP did not send full packet"))
			}
		}
	}
}
