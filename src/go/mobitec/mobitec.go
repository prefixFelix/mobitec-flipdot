package mobitec

import (
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/fonts"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/matrix"
)

var initialized = false

type Display struct {
	m          messenger
	font       fonts.Font
	dataBuffer []packetData
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
		font: fonts.Get("7px"),
	}
}

//
// Display Functions. Any Data is sent to the Display immediately.
//

// Text sends the given text to the display.
// The font used is the one set with Display.SetFont.
func (d *Display) Text(text string) error {
	data := d.m.TextData(text, d.font)
	pkt := d.m.Packet(data)
	return d.m.Send(pkt)
}

// Matrix sends the given matrix to the display.
func (d *Display) Matrix(matrix matrix.Matrix) error {
	data := d.m.MatrixData(matrix)
	pkt := d.m.Packet(data)
	return d.m.Send(pkt)
}

func (d *Display) BufferMatrix(matrix matrix.Matrix) {
	data := d.m.MatrixData(matrix)
	d.dataBuffer = append(d.dataBuffer, data)
}

func (d *Display) BufferText(text string) {
	data := d.m.TextData(text, d.font)
	d.dataBuffer = append(d.dataBuffer, data)
}

func (d *Display) SendBuffer() error {
	data := d.dataBuffer[0]
	print(data)
	for i := 1; i < len(d.dataBuffer); i++ {
		data = append(data, d.dataBuffer[i]...)
	}
	print(data)
	pkt := d.m.Packet(data)
	return d.m.Send(pkt)
}

//
// Setters
//

func (d *Display) SetAddress(address byte) {
	d.m.address = address
}

func (d *Display) SetFont(key string) {
	d.font = fonts.Get(key)
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
