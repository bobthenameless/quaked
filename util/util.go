package util


type Vec3int struct {
	X uint16,
	Y uint16,
	Z uint16
}


// brushes use ints, message 
type Vec3float struct {	
	X float32,
	Y float32,
	Z float32
}


// Brush format, may not be needed by the server depending on how it transmits packets but if it does here it is
// Example brush:
// ( 128 0 0 ) ( 128 1 0 ) ( 128 0 1 ) GROUND1_6 0 0 0 1.0 1.0

type Brush struct {
	FirstPoint	Vec3, // all three points define a plane and the range of values is (128, 128, 64) - (256, 384, 128)
	SecondPoint	Vec3,
	ThirdPoint	Vec3,
	Texture 	string, // does it need it's own type?
	X_Offset 	uint16,
	Y_Offset 	unint16,
	Rotation 	uint16, // degrees 0-359
	X_scale 	float32,
	Y_scale 	float32
}


// I need to read up on the reader/writter buffers in Go but this should be something to enable brushes to be read/written
// not sure what format they're stored in or how (or if) they're sent over the wire but still seemed a likely thing
type BrushReader interface {
	ReadBrush() 	Brush,
	WriteBrush() 	(uint, error) 
}


