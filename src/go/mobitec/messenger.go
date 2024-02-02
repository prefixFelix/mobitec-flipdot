package mobitec

import (
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"time"
)

const (
	DELIMITER = 0xff
)

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
	return []byte{m.address, 0xa2, 0xd0, m.width, 0xd1, m.height}
}
