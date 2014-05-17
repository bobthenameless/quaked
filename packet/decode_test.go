package packet

import (
	"bytes"
	"testing"
)

func TestDecodeConnectionRequest(t *testing.T) {
	packet_bytes := []byte{0x80, 0x00, 0x00, 0x0c, 0x01, 0x51, 0x55, 0x41, 0x4b, 0x45, 0x00, 0x03}
	p, err := Decode(bytes.NewReader(packet_bytes))

	if err != nil {
		t.Error("Failed to read byte array")
	}

	packet, ok := p.(*ConnectionRequestPacket)

	if !ok {
		t.Error("Failed to cast type of packet to *ConnectionRequestPacket")
	}

	if packet.GameName != "QUAKE" {
		t.Errorf("Expected GameName to be 'QUAKE', found '%s' instead", packet.GameName)
	}

	if packet.ProtocolVersion != 3 {
		t.Errorf("Expected ProtocolVersion to be %d, found %d instead", packet.ProtocolVersion)
	}
}

func TestDecodeServerInfoRequest(t *testing.T) {
	packet_bytes := []byte{0x80, 0x00, 0x00, 0x0c, 0x02, 0x51, 0x55, 0x41, 0x4b, 0x45, 0x00, 0x03}
	p, err := Decode(bytes.NewReader(packet_bytes))

	if err != nil {
		t.Error("Failed to read byte array")
	}

	packet, ok := p.(*ServerInfoRequestPacket)

	if !ok {
		t.Error("Failed to cast type of packet to *ServerInfoRequestPacket")
	}

	if packet.GameName != "QUAKE" {
		t.Errorf("Expected GameName to be 'QUAKE', found '%s' instead", packet.GameName)
	}

	if packet.ProtocolVersion != 3 {
		t.Errorf("Expected ProtocolVersion to be %d, found %d instead", packet.ProtocolVersion)
	}
}
