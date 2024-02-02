package mobitec

import (
	"github.com/prefixFelix/mobitec-flipdot/go/mobitec/fonts"
	"github.com/prefixFelix/mobitec-flipdot/go/mobitec/matrix"
)

type Display struct {
	m    messenger
	font fonts.Font
}

// NewDisplay initializes a new Display with the given port, width, and height.
// The boards address is assumed to be 0x06, otherwise it can be set with Display.SetAddress.
func NewDisplay(port string, width, height int) *Display {
	matrix.Width = width
	matrix.Height = height

	return &Display{
		m: messenger{
			address: 0x06,
			port:    port,
			width:   byte(width),
			height:  byte(height),
		},
		font: fonts.FONTS.Get("7px"),
	}
}

//
// Display Functions. Any Data is sent to the Display immediately.
//

// Text sends the given text to the display.
// The font used is the one set with Display.SetFont.
func (d *Display) Text(text string) error {
	data := d.m.dataHeader(d.font.Code, 0, d.font.Height)

	for _, char := range text {
		if mappedChar, ok := fonts.CHARMAP[char]; ok {
			data = append(data, mappedChar)
		} else {
			data = append(data, byte(char))
		}
	}

	return d.m.Send(d.m.Packet(data))
}

func (d *Display) Matrix(matrix matrix.Matrix) error {
	var data []byte
	scm := matrix.ToSubcolumn()
	for band := range scm {
		dataHeader := d.m.dataHeader(0x77, 0, band*5+4)
		data = append(data, dataHeader...)
		for subcolumn := 0; subcolumn < int(d.m.width); subcolumn++ {
			data = append(data, addBits(scm[band][subcolumn]))
		}
	}

	return d.m.Send(d.m.Packet(data))
}

//
// Setters
//

func (d *Display) SetAddress(address byte) {
	d.m.address = address
}

func (d *Display) SetFont(key string) {
	d.font = fonts.FONTS.Get(key)
}

//
// Passthrough
//

// SendSingle opens the serial connection, send the packet, and closes the connection.
func (d *Display) SendSingle(p packet) error {
	return d.m.SendSingle(p)
}

// Open opens the serial connection.
func (d *Display) Open() error {
	return d.m.Open()
}

// Close closes the serial connection.
func (d *Display) Close() error {
	return d.m.Close()
}

// Send sends the packet to the display.
// A serial connection must be opened before calling this method.
func (d *Display) Send(p packet) error {
	return d.m.Send(p)
}

//
// Helpers
//

// addBits takes a slice of bits (as ints) and adds them according
func addBits(bits []bool) byte {
	ret := 32
	for i, bit := range bits {
		ret += asInt(bit) * (1 << i)
	}
	return byte(ret)
}

func asInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
