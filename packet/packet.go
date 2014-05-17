package packet

type buint16 uint16
type buint32 uint32

const (
	ControlMessage         = 0x8000
	ReliableGameMessage    = 0x0001
	ReliableGameMessageEnd = 0x0009
	ReliableGameMessageAck = 0x0002
	UnreliableGameMessage  = 0x0010
)

type PacketHeader struct {
	Type   buint16
	Length buint16
}

const (
	ConnectionRequest        = 0x01
	ConnectionResponseAccept = 0x81
	ConnectionResponseReject = 0x82
	ServerInfoRequest        = 0x02
	ServerInfoResponse       = 0x83
	PlayerInfoRequest        = 0x03
	PlayerInfoResponse       = 0x84
	RuleInfoRequest          = 0x04
	RuleInfoResponse         = 0x85
)

type ControlPacketHeader struct {
	Opcode uint8
}

type ConnectionRequestPacket struct {
	GameName        string
	ProtocolVersion uint8
}

type ConnectionResponseAcceptPacket struct {
	NewPort uint16
	Pad     uint16
}

type ConnectionResponseRejectPacket struct {
	Reason string
}

type ServerInfoRequestPacket struct {
	GameName        string
	ProtocolVersion uint8
}

type ServerInfoResponsePacket struct {
	Address         string
	HostName        string
	LevelName       string
	NumPlayers      uint8
	MaxPlayers      uint8
	ProtocolVersion uint8
}

type PlayerInfoRequestPacket struct {
	PlayerNumber uint8
}

type PlayerInfoResponsePacket struct {
	PlayerNumber uint8
	Name         string
	Colors       uint8
	Pad0         uint8
	Pad1         uint8
	Pad2         uint8
	Frags        uint32
	ConnectTime  uint32
	Address      string
}

type RuleInfoRequestPacket struct {
	Rule string
}

type RuleInfoResponsePacket struct {
	Rule  string
	Value string
}

type GamePacketHeader struct {
	SequenceNumber buint32
}

type GameClientPacketCode uint8

const (
	Nop            = 0x00
	KeepAlive      = 0x01
	Disconnect     = 0x02
	ClientMovement = 0x03
	ConsoleOrder   = 0x04
)

type GameClientPacketHeader struct {
	Opcode GameClientPacketCode
}

type ClientMovementPacket struct {
	ActionTime float32
	TiltAngle  int8
	YawAngle   int8
	FlipAngle  int8
	SpeedFront int16
	SpeedRight int16
	SpeedUp    int16
	Flag       uint8
	Impulse    uint8
}
