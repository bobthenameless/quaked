package packet

import (
	"bytes"
	"testing"
)

func TestConnectionRequestAccept(t *testing.T) {
	packet := &ConnectionResponseAcceptPacket{53101, 0}
	buf := &bytes.Buffer{}
	err := Encode(buf, packet)

	if err != nil {
		t.Errorf("Failed to write byte array, error = %s", err)
	}

	packet_bytes := make([]byte, 4096)

	n, err := buf.Read(packet_bytes)

	if err != nil {
		t.Error("Failed to read buffer")
	}

	if n != 9 {
		t.Errorf("Expected to write 9 bytes, wrote %d instead", n)
	}

	if !bytes.Equal(packet_bytes[:n], []byte{0x80, 0x00, 0x00, 0x09, 0x81, 0x6d, 0xcf, 0x00, 0x00}) {
		t.Errorf("Saw wrong byte sequence, found %d", packet_bytes[:n])
	}
}
