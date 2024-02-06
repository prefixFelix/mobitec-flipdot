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

// TextCustomFont sends the given text to the display using the given custom font.
func (d *Display) TextCustomFont(text string, font fonts.CustomFont) error {
	m := font.GetMatrix(text)
	return d.Matrix(*m)
}

// BufferText adds the given text to the buffer.
// When finished, the buffer can then be sent to the display with Display.SendBuffer.
func (d *Display) BufferText(text string) {
	data := d.m.TextData(text, d.font)
	d.dataBuffer = append(d.dataBuffer, data)
}

// BufferTextAt adds the given text to the buffer at the given offset.
func (d *Display) BufferTextAt(offX, offY int, text string) {
	d.SetOffset(offX, offY)
	d.BufferText(text)
}

// BufferTextWith adds the given text to the buffer using the given font.
func (d *Display) BufferTextWith(font string, text string) {
	d.SetFont(font)
	d.BufferText(text)
}

// BufferTextWithAt adds the given text to the buffer using the given font and offset.
func (d *Display) BufferTextWithAt(font string, offX, offY int, text string) {
	d.SetFont(font)
	d.BufferTextAt(offX, offY, text)
}

// BufferMatrix adds the given matrix to the buffer.
// When finished, the buffer can then be sent to the display with Display.SendBuffer.
func (d *Display) BufferMatrix(matrix matrix.Matrix) {
	data := d.m.MatrixData(matrix)
	d.dataBuffer = append(d.dataBuffer, data)
}

// SendBuffer sends the buffer to the display. All Data is OR'd together, meaning that any dots that are ON in any
// of the buffers will be set in the final display. Everything else is reset.
func (d *Display) SendBuffer() error {
	data := d.dataBuffer[0]
	for i := 1; i < len(d.dataBuffer); i++ {
		data = append(data, d.dataBuffer[i]...)
	}
	pkt := d.m.Packet(data)
	return d.m.Send(pkt)
}

// Fill sends a packet to the display that sets all dots to ON.
func (d *Display) Fill() error {
	m := matrix.NewFullDisplayMatrix()
	return d.Matrix(*m)
}

// Clear sends a packet to the display that sets all dots to OFF.
func (d *Display) Clear() error {
	m := matrix.NewDisplayMatrix()
	return d.Matrix(*m)
}

//
// Setters
//

// SetAddress can be used to set the board's address. 0x6 by default.
// The address is set on the board itself via the switch.
func (d *Display) SetAddress(address byte) {
	d.m.address = address
}

// SetFont sets the font for the display. Does not apply to previously buffered messages.
func (d *Display) SetFont(key string) {
	d.font = fonts.Get(key)
}

// SetOffset sets the horizontal and vertical offset for the text.
// It does not apply to the matrix.
func (d *Display) SetOffset(horizontal, vertical int) {
	d.m.hOffset = horizontal
	d.m.vOffset = vertical
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
