package main

import (
	"machine"
	"time"
)

const (
	ROW_DRIVER_ROWS    = 16
	COL_DRIVER_COLUMNS = 28
)

// DriveDirection represents the direction of current flow
type DriveDirection bool

const (
	DriveHigh DriveDirection = true
	DriveLow  DriveDirection = false
)

// LEDMatrix represents the LED matrix display
type LEDMatrix struct {
	rows    int
	cols    int
	rowAddr [5]machine.Pin
	rowHigh machine.Pin
	rowLow  machine.Pin
	colAddr [5]machine.Pin
	colHL   machine.Pin
	colSel  []machine.Pin

	bufferActive [][]bool
	bufferShadow [][]bool

	rowLookup []uint8
	colLookup []uint8
}

// NewLEDMatrix creates a new LED matrix controller
func NewLEDMatrix(rows, cols int) *LEDMatrix {
	m := &LEDMatrix{
		rows: rows,
		cols: cols,
		// Initialize GPIO pins
		rowAddr: [5]machine.Pin{
			machine.GP2,
			machine.GP3,
			machine.GP4,
			machine.GP5,
			machine.GP6,
		},
		rowHigh: machine.GP7,
		rowLow:  machine.GP8,
		colAddr: [5]machine.Pin{
			machine.GP9,
			machine.GP10,
			machine.GP11,
			machine.GP12,
			machine.GP13,
		},
		colHL:  machine.GP14,
		colSel: []machine.Pin{machine.GP16, machine.GP17},

		// Initialize lookup tables
		rowLookup: []uint8{1, 4, 5, 2, 3, 6, 7, 9, 12, 13, 10, 11, 14, 15, 17, 20},
		colLookup: []uint8{
			1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18,
			19, 20, 21, 22, 23, 25, 26, 27, 28, 29, 30, 31,
		},
	}

	// Initialize display buffers
	m.bufferActive = make([][]bool, rows)
	m.bufferShadow = make([][]bool, rows)
	for i := range m.bufferActive {
		m.bufferActive[i] = make([]bool, cols)
		m.bufferShadow[i] = make([]bool, cols)
	}

	// Configure all pins as outputs
	for i := range m.rowAddr {
		m.rowAddr[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	m.rowHigh.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.rowLow.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := range m.colAddr {
		m.colAddr[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	m.colHL.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := range m.colSel {
		m.colSel[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	return m
}

// setRowAddress sets the row address pins based on lookup table
func (m *LEDMatrix) setRowAddress(row uint8) {
	addr := m.rowLookup[row]
	for i := range m.rowAddr {
		if (addr & (1 << uint8(i))) != 0 {
			m.rowAddr[i].High()
		} else {
			m.rowAddr[i].Low()
		}
	}
}

// setRowDirection sets the row direction (high or low drive)
func (m *LEDMatrix) setRowDirection(dir DriveDirection) {
	if dir == DriveHigh {
		m.rowLow.Low()
		m.rowHigh.High()
	} else {
		m.rowHigh.Low()
		m.rowLow.High()
	}
}

// setColumnAddress sets the column address pins based on lookup table
func (m *LEDMatrix) setColumnAddress(col uint8) {
	addr := m.colLookup[col]
	for i := range m.colAddr {
		if (addr & (1 << uint8(i))) != 0 {
			m.colAddr[i].High()
		} else {
			m.colAddr[i].Low()
		}
	}
}

// setColumnDirection sets the column direction (high or low drive)
func (m *LEDMatrix) setColumnDirection(dir DriveDirection) {
	if dir == DriveHigh {
		m.colHL.Low()
	} else {
		m.colHL.High()
	}
}

// setColumnEnabled enables or disables column drivers
func (m *LEDMatrix) setColumnEnabled(col int) {
	for i := range m.colSel {
		m.colSel[i].Low()
	}
	if col >= 0 && col < len(m.colSel) {
		m.colSel[col].High()
	}
}

// drivePixel drives a single pixel
func (m *LEDMatrix) drivePixel(row, col int, dir DriveDirection) {
	m.setRowAddress(uint8(row))
	m.setColumnAddress(uint8(col % COL_DRIVER_COLUMNS))

	m.setRowDirection(dir)
	m.setColumnDirection(!dir)

	m.setColumnEnabled(col / COL_DRIVER_COLUMNS)
	time.Sleep(time.Microsecond * 190)
	m.setColumnEnabled(-1) // Disable all columns
}

// Clear clears the display buffer
func (m *LEDMatrix) Clear() {
	for i := range m.bufferActive {
		for j := range m.bufferActive[i] {
			m.bufferActive[i][j] = false
		}
	}
}

// Fill fills the display buffer with a value
func (m *LEDMatrix) Fill(value bool) {
	for i := range m.bufferActive {
		for j := range m.bufferActive[i] {
			m.bufferActive[i][j] = value
		}
	}
}

// SetPixel sets a pixel in the display buffer
func (m *LEDMatrix) SetPixel(row, col int, value bool) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.bufferActive[row][col] = value
	}
}

// Refresh refreshes the display, updating changed pixels
func (m *LEDMatrix) Refresh(forceRefresh bool) {
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			if forceRefresh || m.bufferActive[row][col] != m.bufferShadow[row][col] {
				dir := DriveHigh
				if m.bufferActive[row][col] {
					dir = DriveLow
				}
				m.drivePixel(row, col, dir)
				m.bufferShadow[row][col] = m.bufferActive[row][col]
				time.Sleep(time.Microsecond * 10)
			}
		}
	}
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	display := NewLEDMatrix(ROW_DRIVER_ROWS, COL_DRIVER_COLUMNS)

	// Initial test pattern
	led.High()

	// Clear and deep refresh
	display.Clear()
	display.Refresh(true)
	time.Sleep(time.Millisecond * 2000)

	// Blink pattern
	for i := 0; i < 3; i++ {
		display.Fill(true)
		display.Refresh(false)
		time.Sleep(time.Millisecond * 50)

		display.Fill(false)
		display.Refresh(false)
		time.Sleep(time.Millisecond * 50)
	}

	led.Low()

	// Main loop
	for {
		time.Sleep(time.Millisecond * 100)
	}
}
