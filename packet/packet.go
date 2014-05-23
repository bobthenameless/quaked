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

//should let us switch on *UnreliableMessagePacket within Encode()
type UnreliableMessagePacket interface {
	Header	PacketHeader	
	SequenceNumber	GamePacketHeader
	Messages	[]uint8 // size is header.length-sizeof(SequenceNum)
	AddOpcode(uint8) error
	
	
}


type GameServerPacketCode uint8

const (
	Sv_Nop            = 0x00
	Sv_KeepAlive      = 0x01
	Sv_Disconnect     = 0x02
	Sv_ClientUpdate	  = 0x03
	//04
	//05
	//06
	Sv_TimeStamp	  = 0x07

	Sv_ClientData	  = 0x0F
)

type GameServerPacket struct {
	Header	PacketHeader	
	SequenceNumber	GamePacketHeader
	Messages	[]uint8 // size is header.length-sizeof(SequenceNum)
	
}

// GameServerPacket now implements (UnreliableMessagePacket) interface, and so cast it and have it trigger the *UnreliableMessagePacket switch case
func (p GameServerPacket) AddOpcode(opcode GameServerPacketCode) error {
	switch opcode {
	default:
		err = fmt.Errorf("Unrecognized server message opcode %d", opcode)
	case Sv_TimeStamp:
		time float32 = 0.0f // time in secs since start of level
		new_timestamp = &Sv_TimeStampPacket{0x07, time}
		p.Messages.append(new_timestap) // not real Go

	case Sv_ClientData:
		info_mask = 0x0200 // update client inventory packet sent
		client_update_p = &Sv_ClientDataPacket{}.Build(info_mask)

		// client_update_p should be == to 
		// << 0f 02 00 00 00 10 21 00 64 64 64 00 0a 00 01 >>
		//    i believe
		// for: opcode = 07
		//  	mask == 0x0200
		//	inventory_mask = 0x00001021
		//	health == 0x0064  
		//	current_ammo == 0x64
		//	ammo_shells == 0x64
		//	ammo_nails == 0x00
		//	ammo_rockets == 0x0a 
		//	ammo_cells == 0x00
		//	weapon == 0x01
		// this should give the client's status:
		//health: 100 shells: 100  rockets: 10
		//inv: shotgun, axe, rocket launcher
	// set new p.Header.Length = len(Messages)
	}
	
}

type Sv_TimeStampPacket struct {
	Opcode	uint8 // 0x07
	Time	float32
}

//const decl Sv_ClientDataMask
const (
	Sv_ClientDataMask_ViewOfsZ uint16 = 1 << iota	
	Sv_ClientDataMask_AngOfs_one // 0x2

	Sv_ClientDataMask_Angles_zero 
	Sv_ClientDataMask_Angles_one
	Sv_ClientDataMask_Angles_two //0x10

	Sv_ClientDataMask_Vel_zero
	Sv_ClientDataMask_Vel_one
	Sv_ClientDataMask_Vel_two

	Sv_ClientDataMask_Inventory_Mask // 0x0200
	Sv_ClientDataMask_Unused1
	Sv_ClientDataMask_Unused2

	Sv_ClientDataMask_WeaponFrame // 0x1000
	Sv_ClientDataMask_ArmorValue 
	Sv_ClientDataMask_WeaponModel // 0x4000
)


type Sv_ClientDataPacket struct {
	Opcode	uint8
	Mask	uint16 // following line comments show mask values, see const decl Sv_ClientDataMask
	View_Ofs_Z float32 // & 0x0001
	Ang_ofs_1 float32 // & 0x0002
	Angles	[3]uint8 // vec3; 0x0004, & 0x0008, 0x0010
	Vel	[3]uint8 // vec3; 0x0020, 0x0040, 0x0080  
	Items 	uint32 // inventory mask; 0x0200 ; const_inventory_mask 
	WeaponFrame	uint8 // 0x1000
	ArmorValue	uint8 // 0x2000
	WeaponModel	uint8 // 0x4000, end of mask.
	Health		uint16 // always present, same for rest of struct
	CurrentAmmo	uint8
	Ammo_Shells	uint8
	Ammo_Nails	uint8
	Ammo_Rockets	uint8
	Ammo_Cells	uint8
	Weapon		uint8	
}

// theoretical future usage
// mask = 0x0064 = Sv_ClientDataMask_Angles_zero | Sv_CDM_vel0 | Sv_CDM_vel1
// p = Sv_ClientDataPacket{}.Build(mask)
// p == {0x0f, mask,        //always present
// 		angles[0] uint8, vel[0] uint8, vel[1] uint8,          //present because mask said so
//  Health, CurrentAmmo, Ammo_shells, AmmoNails, Ammo_Rockets, Ammo_Cells, Weapon}   // always present


//non working outline
func (p Sv_ClientDataPacket) Build(mask uint16) []uint8 {
	out_buf []uint8 := make() // read up on slices/buffers/whatev notation, buffers are what i probably want
	out_buf.append(0x07)
	out_buf.append(mask)
	
	/*
	if mask & 0x0001
		append view_ofs_z
	...
	if mask & 0x0010
		append angles[2]
	..
	if mask & 0x1000
		append weaponframe
	
	*/

	//Sv_ClientDataMask_Inventory_Mask
	if mask & 0200 != 0 {

		current_inventory = 0x1001 // shotgun and axe only

		it_rocket_launcher = 0x00000020
		
		//give cleint rocket launcher in inventory		
		out_buf.append(current_inventory | it_rocket_launcher)
		
	}
	
	out_buf.append(p.health)
	out_buf.append(p.currentammo)
	out_buf.append(p.ammo_shells)
	out_buf.append(p.ammo_nails)
	out_buf.append(p.ammo_rockets)
	out_buf.append(p.ammo_cells)
	out_buf.append(p.weapon)

	return outbuf
}
