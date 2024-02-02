package mobitec

import (
	"errors"
	"fmt"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/fonts"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/matrix"
	"github.com/tarm/serial"
	"time"
)

const (
	DELIMITER = 0xff
)

type packetData []byte

type packet []byte

func (p *packet) Add(b ...byte) {
	*p = append(*p, b...)
}

// calculates packet checksum
func (p *packet) checksum() []byte {
	var packetSum int
	for _, b := range (*p)[1:] {
		packetSum += int(b)
	}
	packetSum &= 0xff

	switch packetSum {
	case 0xfe:
		return []byte{0xfe, 0x00}
	case 0xff:
		return []byte{0xfe, 0x01}
	default:
		return []byte{byte(packetSum)}
	}
}

type messenger struct {
	address byte
	port    string
	width   byte
	height  byte
	serial  *serial.Port
}

// Packet returns a mobitec packet with the given data.
func (m *messenger) Packet(data []byte) packet {
	p := packet{DELIMITER}
	p.Add(m.packetHeader()...)
	p.Add(data...)
	p.Add(p.checksum()...)
	p.Add(DELIMITER)
	return p
}

// TextData returns a mobitec packet with the given text and font.
func (m *messenger) TextData(text string, font fonts.Font) packetData {
	data := m.dataHeader(font.Code, 0, font.Height)

	for _, char := range text {
		if mappedChar, ok := fonts.CHARMAP[char]; ok {
			data = append(data, mappedChar)
		} else {
			data = append(data, byte(char))
		}
	}

	return data
}

func (m *messenger) MatrixData(matrix matrix.Matrix) packetData {
	var data []byte
	scm := matrix.ToSubcolumn()
	for band := range scm {
		dataHeader := m.dataHeader(fonts.Get(fonts.Font_pixel_subcolumns).Code, 0, band*5+4)
		data = append(data, dataHeader...)
		for subcolumn := 0; subcolumn < int(m.width); subcolumn++ {
			data = append(data, addBits(scm[band][subcolumn]))
		}
	}

	return data
}

//
// Serial
//

func (m *messenger) Open() (err error) {
	c := &serial.Config{Name: m.port, Baud: 4800, ReadTimeout: time.Second}
	m.serial, err = serial.OpenPort(c)
	if err != nil {
		return fmt.Errorf("failed to open serial port: %v", err)
	}

	return nil
}

func (m *messenger) Close() error {
	return m.serial.Close()
}

func (m *messenger) Send(p packet) error {
	if m.serial == nil {
		return errors.New("serial port not open")
	}

	_, err := m.serial.Write(p)
	if err != nil {
		return fmt.Errorf("failed to write to serial port: %v", err)
	}
	return nil
}

func (m *messenger) SendSingle(p packet) error {
	err := m.Open()
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, m.serial.Close())
	}()
	err = m.Send(p)

	return err
}

//
// Packet Generation
//

func (m *messenger) dataHeader(font byte, hOffset, vOffset int) []byte {
	return []byte{0xd2, byte(hOffset), 0xd3, byte(vOffset), 0xd4, font}
}

// returns the packet header
func (m *messenger) packetHeader() []byte {
	data := []byte{m.address, 0xa2}
	if !initialized {
		data = append(data, 0xd0, m.width, 0xd1, m.height)
		initialized = true
	}

	return data
}

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
