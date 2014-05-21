package message

import (
	"github.com/bobthenameless/quaked/util"
)

// i didn't reference or really look at how these are implemented in decode/encode.go
// main source of message info http://www.gamers.org/dEngine/quake/Qdem/dem-1.0.2-5.html
// server broadcast use the same thing as .DEM file format

type MessageBlock struct {
	BlockSize 	uint32
	Angles 		[3]float32
	Messages 	[]Message // will be an array of Messages, length is 
				  // detirmined by BlockSize. not sure if this
				  // works or a better go idiom exists
	
}

type messageid uint16

type Message interface {
	ID messageid
	// MessageContents string
}

const (
	BadMessage = 0x00
	NoopMessage = iota
	DisconnectMessage = iota
	UpdatePlayerStateMessage = iota
	VersionMessage = iota		// 0x04	
	SetViewMessage = iota
	SoundMessage = iota
)

// The Following was an attempt to do things wrong but might be able to re-appropriate
/*
type BadMessage struct {
	ID messageid
}

//Keep alive message
type NoopMessage struct {
	ID = 0x01 // argh this is illegal and i need to learn more about go before i continue
}

type DisconnectMessage struct {
	ID = 0x02
//disconnect from server
}

type UpdatePlayerStateMessage struct {
	ID 		= 0x03
	Index 		uint32 // index into playerstate array
	Value 		uint32
	PlayerState 	[32]uint32
}
*/

/*-----  possible indices for playerstate:

0 health	1 n/a		2 weaponmodel	3 currentammo
4 armorvalue	5 weaponframe	6 ammo_shells	7 ammo_nails
8 ammo_rockets	9 ammo_cells   10 weapon       11 total_secrets
12 total_monsters 13 found_secrets 14 found_monsters

------*/

/*
type VersionMessage struct {
	ID		= 0x04
	ServerVersion	uint32 // "should be 0x0F", see also: serverinfo message
}

type SetViewMessage struct {
	ID		= 0x05
	Entity		uint32 		// camera entity
}

*/

/* A list of messages for the .DEM format

    6.7 sound
    6.8 time
    6.9 print
    6.10 stufftext
    6.11 setangle
    6.12 serverinfo
    6.13 lightstyle
    6.14 updatename
    6.15 updatefrags
    6.16 clientdata
    6.17 stopsound
    6.18 updatecolors
    6.19 particle
    6.20 damage
    6.21 spawnstatic
    6.22 spawnbinary
    6.23 spawnbaseline
    6.24 temp_entity
    6.25 setpause
    6.26 signonum
    6.27 centerprint
    6.28 killedmoster
    6.29 foundsecret
    6.30 spawnstaticsound
    6.31 intermission
    6.32 finale
    6.33 cdtrack
    6.34 sellscreen
    6.35 updateentity 

*/
