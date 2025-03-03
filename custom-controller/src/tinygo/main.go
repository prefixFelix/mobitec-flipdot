//
// Based on flipdot-games Rust code from Anton Berneving (https://github.com/antbern/flipdot-games/tree/main/pico-firmware/src/driver.rs)
//

package main

import (
	"machine"
	"time"
	_ "embed"
)

// DriveDirection represents the direction of current flow
type DriveDirection bool

const (
	ROW_DRIVER_ROWS    = 16
	COL_DRIVER_COLUMNS = 126
	DriveHigh DriveDirection = true
	DriveLow  DriveDirection = false
)

//go:embed fonts/dot01.bin
var dot01 []byte
//go:embed fonts/dot02.bin
var dot02 []byte
//go:embed fonts/dot03.bin
var dot03 []byte
//go:embed fonts/dot04.bin
var dot04 []byte
//go:embed fonts/dot05.bin
var dot05 []byte
//go:embed fonts/dot06.bin
var dot06 []byte
//go:embed fonts/dot06Num.bin
var dot06Num []byte
//go:embed fonts/dot07.bin
var dot07 []byte
//go:embed fonts/dot08.bin
var dot08 []byte
//go:embed fonts/dot09.bin
var dot09 []byte
//go:embed fonts/dot10.bin
var dot10 []byte
//go:embed fonts/dot11.bin
var dot11 []byte
//go:embed fonts/ww4x7.bin
var ww4x7 []byte
//go:embed fonts/ww4x9.bin
var ww4x9 []byte
//go:embed fonts/ww5x7.bin
var ww5x7 []byte
//go:embed fonts/ww5x9.bin
var ww5x9 []byte
//go:embed fonts/wwXx7.bin
var wwXx7 []byte
//go:embed fonts/dotCousineBold.bin
var dotCousineBold []byte
//go:embed fonts/dotDejaVuBold.bin
var dotDejaVuBold []byte
//go:embed fonts/dotSansRegular.bin
var dotSansRegular []byte
//go:embed fonts/bus.bin
var bus []byte
//go:embed fonts/moon.bin
var moon []byte

// Font represents a bitmap font for the flipdot display
type Font struct {
    Width       uint8    // Font width in pixels
    Height      uint8    // Font height in pixels
    FirstChar   uint8    // First character in the font set (usually ASCII value)
    CharCount   uint8    // Number of characters in the font set
    JumpTable   []JumpTableEntry
    BitmapData  []byte   // The actual bitmap data
}

// JumpTableEntry contains information for quick access to character data
type JumpTableEntry struct {
    JumpAddress uint16 // Offset to this character's data in BitmapData
    Size        uint8  // Size in bytes
    Width       uint8  // Width in pixels
}

// FlipdotMatrix represents the Flipdot matrix display
type FlipdotMatrix struct {
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

	currentFont *Font
}

// NewFlipdotMatrix creates a new Flipdot matrix controller
func NewFlipdotMatrix(cols, rows int) *FlipdotMatrix {
	m := &FlipdotMatrix{
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
		colSel: []machine.Pin{
            machine.GP16,  // 28 cols matrix
            machine.GP17,  // 56 cols matrix
            machine.GP18,  // 84 cols matrix
            machine.GP19,  // 112 cols matrix
            machine.GP20,  // 140 cols matrix
        },

		// Initialize lookup tables
		rowLookup: []uint8{1, 4, 5, 2, 3, 6, 7, 9, 12, 13, 10, 11, 14, 15, 17, 20},
		colLookup: []uint8{
			1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 41, 42, 43, 44, 45, 46, 47, 49, 50, 51, 52, 53, 54, 55, 57, 58, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 73, 74, 75, 76, 77, 78, 79, 81, 82, 83, 84, 85, 86, 87, 89, 90, 91, 92, 93, 94, 95, 97, 98, 99, 100, 101, 102, 103, 105, 106, 107, 108, 109, 110, 111, 113, 114, 115, 116, 117, 118, 119, 121, 122, 123, 124, 125, 126, 127, 129, 130, 131, 132, 133, 134, 135, 137, 138, 139, 140, 141, 142, 143, 145, 146, 147, 148, 149, 150, 151, 153, 154, 155, 156, 157, 158, 159,
		},
		currentFont: nil,
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
func (m *FlipdotMatrix) setRowAddress(row uint8) {
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
func (m *FlipdotMatrix) setRowDirection(dir DriveDirection) {
	if dir == DriveHigh {
		m.rowLow.Low()
		m.rowHigh.High()
	} else {
		m.rowHigh.Low()
		m.rowLow.High()
	}
}

// setColumnAddress sets the column address pins based on lookup table
func (m *FlipdotMatrix) setColumnAddress(col uint8) {
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
func (m *FlipdotMatrix) setColumnDirection(dir DriveDirection) {
	if dir == DriveHigh {
		m.colHL.Low()
	} else {
		m.colHL.High()
	}
}

// setColumnEnabled enables or disables column drivers
func (m *FlipdotMatrix) setColumnEnabled(col int) {
	for i := range m.colSel {
		m.colSel[i].Low()
	}
	if col >= 0 && col < len(m.colSel) {
		m.colSel[col].High()
	}
}

// drivePixel drives a single pixel
func (m *FlipdotMatrix) drivePixel(col, row int, dir DriveDirection) {
	m.setRowAddress(uint8(row))
	m.setColumnAddress(uint8(col % COL_DRIVER_COLUMNS))

	m.setRowDirection(dir)
	m.setColumnDirection(!dir)

	m.setColumnEnabled(col / 28)
	time.Sleep(time.Microsecond * 200)
	m.setColumnEnabled(-1) // Disable all columns
}

// Refresh refreshes the display, updating changed pixels
func (m *FlipdotMatrix) Refresh(forceRefresh bool) {
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			if forceRefresh || m.bufferActive[row][col] != m.bufferShadow[row][col] {
				dir := DriveHigh
				if m.bufferActive[row][col] {
					dir = DriveLow
				}
				m.drivePixel(col, row, dir)
				m.bufferShadow[row][col] = m.bufferActive[row][col]
				time.Sleep(time.Microsecond * 10)
			}
		}
	}
}

// Clear clears the display buffer
func (m *FlipdotMatrix) Clear() {
	m.Fill(false)
}

// Fill fills the display buffer with a value
func (m *FlipdotMatrix) Fill(value bool) {
	for i := range m.bufferActive {
		for j := range m.bufferActive[i] {
			m.bufferActive[i][j] = value
		}
	}
}

// SetPixel sets a pixel in the display buffer
func (m *FlipdotMatrix) SetPixel(col, row int, value bool) {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.bufferActive[row][col] = value
	}
}

// Invert toggles the state of all pixels in the active buffer
func (m *FlipdotMatrix) Invert() {
    for row := 0; row < m.rows; row++ {
        for col := 0; col < m.cols; col++ {
            // Toggle each pixel's state in the active buffer
            m.bufferActive[row][col] = !m.bufferActive[row][col]
        }
    }
}

// ShiftHorizontal shifts the active buffer horizontally
// Parameters:
//   steps: Number of steps to shift (positive for right, negative for left)
//   rotate: If true, wraps pixels around; if false, fills with blank pixels
func (m *FlipdotMatrix) ShiftHorizontal(steps int, rotate bool) {
    // No shift needed
    if steps == 0 {
        return
    }

    // Get absolute value and direction
    abs := steps
    if abs < 0 {
        abs = -abs
    }
    shiftRight := steps > 0

    // Don't shift more than matrix width
    if abs > m.cols {
        if rotate {
            // For rotation, take modulo to get effective shift
            abs = abs % m.cols
            if abs == 0 {
                return // Full rotation results in original state
            }
            if steps < 0 {
                abs = m.cols - abs // Convert left shift to equivalent right shift
                shiftRight = true
            }
        } else {
            // For non-rotating shift, capping at matrix width
            abs = m.cols
        }
    }

    // Create a temporary buffer
    temp := make([][]bool, m.rows)
    for i := range temp {
        temp[i] = make([]bool, m.cols)
        for j := range temp[i] {
            temp[i][j] = false
        }
    }

    // Perform the shift
    for row := 0; row < m.rows; row++ {
        for col := 0; col < m.cols; col++ {
            var newCol int
            if shiftRight {
                newCol = col + abs
                if newCol >= m.cols {
                    if rotate {
                        newCol = newCol % m.cols
                    } else {
                        continue // Skip if outside bounds
                    }
                }
                temp[row][newCol] = m.bufferActive[row][col]
            } else {
                newCol = col - abs
                if newCol < 0 {
                    if rotate {
                        newCol = m.cols + newCol // Wrap around
                    } else {
                        continue // Skip if outside bounds
                    }
                }
                temp[row][newCol] = m.bufferActive[row][col]
            }
        }
    }

    // Update the active buffer
    for row := 0; row < m.rows; row++ {
        for col := 0; col < m.cols; col++ {
            m.bufferActive[row][col] = temp[row][col]
        }
    }
}

// ShiftVertical shifts the active buffer vertically
// Parameters:
//   steps: Number of steps to shift (positive for down, negative for up)
//   rotate: If true, wraps pixels around; if false, fills with blank pixels
func (m *FlipdotMatrix) ShiftVertical(steps int, rotate bool) {
    // No shift needed
    if steps == 0 {
        return
    }

    // Get absolute value and direction
    abs := steps
    if abs < 0 {
        abs = -abs
    }
    shiftDown := steps > 0

    // Don't shift more than matrix height
    if abs > m.rows {
        if rotate {
            // For rotation, take modulo to get effective shift
            abs = abs % m.rows
            if abs == 0 {
                return // Full rotation results in original state
            }
            if steps < 0 {
                abs = m.rows - abs // Convert upward shift to equivalent downward shift
                shiftDown = true
            }
        } else {
            // For non-rotating shift, capping at matrix height
            abs = m.rows
        }
    }

    // Create a temporary buffer
    temp := make([][]bool, m.rows)
    for i := range temp {
        temp[i] = make([]bool, m.cols)
        for j := range temp[i] {
            temp[i][j] = false
        }
    }

    // Perform the shift
    for row := 0; row < m.rows; row++ {
        for col := 0; col < m.cols; col++ {
            var newRow int
            if shiftDown {
                newRow = row + abs
                if newRow >= m.rows {
                    if rotate {
                        newRow = newRow % m.rows
                    } else {
                        continue // Skip if outside bounds
                    }
                }
                temp[newRow][col] = m.bufferActive[row][col]
            } else {
                newRow = row - abs
                if newRow < 0 {
                    if rotate {
                        newRow = m.rows + newRow // Wrap around
                    } else {
                        continue // Skip if outside bounds
                    }
                }
                temp[newRow][col] = m.bufferActive[row][col]
            }
        }
    }

    // Update the active buffer
    for row := 0; row < m.rows; row++ {
        for col := 0; col < m.cols; col++ {
            m.bufferActive[row][col] = temp[row][col]
        }
    }
}

// DrawHorizontalLine draws a horizontal line starting at (col, row) with given length
// If the line extends beyond matrix boundaries, it will be clipped
func (m *FlipdotMatrix) DrawHorizontalLine(col, row, length int, value bool) {
    // Skip if row is outside the matrix
    if row < 0 || row >= m.rows {
        return
    }

    // Clip col if negative
    if col < 0 {
        length += col // Reduce length
        col = 0       // Start at first column
    }

    // Return if no length left after clipping
    if length <= 0 {
        return
    }

    // Calculate end point, clipping to matrix width
    endCol := col + length
    if endCol > m.cols {
        endCol = m.cols
    }

    // Draw the line
    for col := col; col < endCol; col++ {
        m.bufferActive[row][col] = value
    }
}

// DrawVerticalLine draws a vertical line starting at (col, row) with given length
// If the line extends beyond matrix boundaries, it will be clipped
func (m *FlipdotMatrix) DrawVerticalLine(col, row, length int, value bool) {
    // Skip if column is outside the matrix
    if col < 0 || col >= m.cols {
        return
    }

    // Clip row if negative
    if row < 0 {
        length += row // Reduce length
        row = 0       // Start at first row
    }

    // Return if no length left after clipping
    if length <= 0 {
        return
    }

    // Calculate end point, clipping to matrix height
    endRow := row + length
    if endRow > m.rows {
        endRow = m.rows
    }

    // Draw the line
    for row := row; row < endRow; row++ {
        m.bufferActive[row][col] = value
    }
}

// DrawRectangle draws a rectangle with top-left corner at (col, row) with specified width and height
// If the rectangle extends beyond matrix boundaries, it will be clipped
func (m *FlipdotMatrix) DrawRectangle(col, row, width, height int, value bool) {
    // Draw horizontal lines (top and bottom)
    m.DrawHorizontalLine(col, row, width, value)
    m.DrawHorizontalLine(col, row+height-1, width, value)

    // Draw vertical lines (left and right)
    m.DrawVerticalLine(col, row, height, value)
    m.DrawVerticalLine(col+width-1, row, height, value)
}

// DrawSquare draws a square with top-left corner at (col, row) with specified side length
// This is a convenience function that calls DrawRectangle with equal width and height
func (m *FlipdotMatrix) DrawSquare(col, row, size int, value bool) {
    m.DrawRectangle(col, row, size, size, value)
}

// TODO NOT TESTED!
// PlaceBuffer places a smaller buffer into the active buffer at a specific position
// Parameters:
//   startRow, startCol: Top-left coordinates where to place the buffer
//   buffer: 2D array of boolean values representing the buffer to place
//   transparent: If true, only "true" values from the source buffer will be copied
func (m *FlipdotMatrix) PlaceBuffer(startCol, startRow int, buffer [][]bool, transparent bool) {
    if buffer == nil || len(buffer) == 0 {
        return
    }

    bufferHeight := len(buffer)
    bufferWidth := len(buffer[0])

    // Determine overlapping region
    startRowClipped := startRow
    if startRowClipped < 0 {
        startRowClipped = 0
    }

    startColClipped := startCol
    if startColClipped < 0 {
        startColClipped = 0
    }

    endRow := startRow + bufferHeight
    if endRow > m.rows {
        endRow = m.rows
    }

    endCol := startCol + bufferWidth
    if endCol > m.cols {
        endCol = m.cols
    }

    // Copy buffer to active buffer, respecting transparency
    for row := startRowClipped; row < endRow; row++ {
        for col := startColClipped; col < endCol; col++ {
            // Calculate source buffer coordinates
            srcRow := row - startRow
            srcCol := col - startCol

            // Skip if source coordinates are out of bounds
            if srcRow < 0 || srcRow >= bufferHeight || srcCol < 0 || srcCol >= bufferWidth {
                continue
            }

            // Apply the pixel from the source buffer
            if !transparent || buffer[srcRow][srcCol] {
                m.bufferActive[row][col] = buffer[srcRow][srcCol]
            }
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

// SetFont sets the font to be used for text rendering
func (m *FlipdotMatrix) SetFont(fontName string) {
    switch fontName {
    case "dot01":
        m.currentFont, _ = ParseFont(dot01)
    case "dot02":
        m.currentFont, _ = ParseFont(dot02)
    case "dot03":
        m.currentFont, _ = ParseFont(dot03)
    case "dot04":
        m.currentFont, _ = ParseFont(dot04)
    case "dot05":
        m.currentFont, _ = ParseFont(dot05)
    case "dot06":
        m.currentFont, _ = ParseFont(dot06)
    case "dot06Num":
        m.currentFont, _ = ParseFont(dot06Num)
    case "dot07":
        m.currentFont, _ = ParseFont(dot07)
    case "dot08":
        m.currentFont, _ = ParseFont(dot08)
    case "dot09":
        m.currentFont, _ = ParseFont(dot09)
    case "dot10":
        m.currentFont, _ = ParseFont(dot10)
    case "dot11":
        m.currentFont, _ = ParseFont(dot11)
    case "ww4x7":
        m.currentFont, _ = ParseFont(ww4x7)
    case "ww4x9":
        m.currentFont, _ = ParseFont(ww4x9)
    case "ww5x7":
        m.currentFont, _ = ParseFont(ww5x7)
    case "ww5x9":
        m.currentFont, _ = ParseFont(ww5x9)
    case "wwXx7":
        m.currentFont, _ = ParseFont(wwXx7)
    case "dotCousineBold":
        m.currentFont, _ = ParseFont(dotCousineBold)
    case "dotDejaVuBold":
        m.currentFont, _ = ParseFont(dotDejaVuBold)
    case "dotSansRegular":
        m.currentFont, _ = ParseFont(dotSansRegular)
    case "moon":
        m.currentFont, _ = ParseFont(moon)
    case "bus":
        m.currentFont, _ = ParseFont(bus)
    }
}

// ParseFont parses a font file in the format similar to what's used in wwFlipGFX
func ParseFont(fontData []byte) (*Font, error) {
    if len(fontData) < 4 {
        return nil, nil
    }

    font := &Font{
        Width:      fontData[0],
        Height:     fontData[1],
        FirstChar:  fontData[2],
        CharCount:  fontData[3],
    }

    // Constants for parsing the jump table
    const (
        jumpTableBytes  = 4  // Bytes per entry in jump table
        jumpTableStart  = 4  // Offset where jump table starts
    )

    // Calculate size of jump table
    jumpTableSize := int(font.CharCount) * jumpTableBytes

    // Ensure font data is long enough
    if len(fontData) < jumpTableStart + jumpTableSize {
        return nil, nil
    }

    // Parse jump table
    font.JumpTable = make([]JumpTableEntry, font.CharCount)
    for i := uint8(0); i < font.CharCount; i++ {
        offset := jumpTableStart + int(i)*jumpTableBytes

        // Parse MSB and LSB to get jump address
        msb := fontData[offset]
        lsb := fontData[offset+1]

        // Special case for characters without bitmap data
        if msb == 0xFF && lsb == 0xFF {
            font.JumpTable[i] = JumpTableEntry{
                JumpAddress: 0xFFFF, // Invalid address
                Size:        fontData[offset+2],
                Width:       fontData[offset+3],
            }
        } else {
            font.JumpTable[i] = JumpTableEntry{
                JumpAddress: uint16(msb)<<8 | uint16(lsb),
                Size:        fontData[offset+2],
                Width:       fontData[offset+3],
            }
        }
    }

    // Calculate bitmap data
    bitmapStart := jumpTableStart + jumpTableSize
    font.BitmapData = fontData[bitmapStart:]

    return font, nil
}

// DrawText renders text at the specified position on the flipdot matrix
func (m *FlipdotMatrix) DrawText(x, y int, text string) error {
    if m.currentFont == nil {
        return nil
    }

    font := m.currentFont
    cursorX := x

    // Process each character in the text
    for _, char := range text {
        // Skip if character is outside the font range
        if char < rune(font.FirstChar) || char >= rune(font.FirstChar+font.CharCount) {
            // You could render a replacement character or just skip
            cursorX += int(font.Width) // Use default width for missing chars
            continue
        }

        // Get character index and jump table entry
        charIndex := uint8(char) - font.FirstChar
        entry := font.JumpTable[charIndex]

        // Skip unmapped characters
        if entry.JumpAddress == 0xFFFF {
            cursorX += int(entry.Width)
            continue
        }

        // Calculate how many bytes per column based on font height
        bytesPerColumn := (int(font.Height) + 7) / 8

        // Draw the character bitmap
        for col := 0; col < int(entry.Width); col++ {
            for byteIndex := 0; byteIndex < bytesPerColumn; byteIndex++ {
                // Calculate the offset in bitmap data
                offset := int(entry.JumpAddress) + col*bytesPerColumn + byteIndex

                // Safety check
                if offset >= len(font.BitmapData) {
                    continue
                }

                // Get the byte containing 8 vertical pixels
                dataByte := font.BitmapData[offset]

                // Draw each bit as a pixel
                for bit := 0; bit < 8; bit++ {
                    rowPosition := y + byteIndex*8 + bit

                    // Skip if we've exceeded the character's height
                    if byteIndex*8+bit >= int(font.Height) {
                        break
                    }

                    // Set pixel if the bit is set
                    if (dataByte & (1 << bit)) != 0 {
                        m.SetPixel(cursorX+col, rowPosition, true)
                    }
                }
            }
        }

        // Move cursor to next character position
        cursorX += int(entry.Width) + 1 // +1 for character spacing
    }

    return nil
}


func main() {
    // TODO: Investigate first dot fonts (with "Test")
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	display := NewFlipdotMatrix(COL_DRIVER_COLUMNS, ROW_DRIVER_ROWS)


	// Clear and deep refresh
	led.High()
	display.Clear()
	display.Refresh(true)
	time.Sleep(time.Millisecond * 2000)
    display.SetPixel(0, 0, true)
    display.SetPixel(1, 0, true)
	display.SetPixel(125, 0, true)
	display.SetPixel(120, 0, true)
	display.Refresh(false)

//     for i := 0; i < COL_DRIVER_COLUMNS; i++ {
//         led.High()
// 		display.SetPixel(i, 0, true)
// 		display.Refresh(false)
// 		time.Sleep(time.Millisecond * 200)
//         led.Low()
// 	}

//     display.SetFont("moon")
//     display.DrawText(0, 0, "(")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("ww4x7")
//     display.DrawText(10, 0, "+")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot01")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()



//     display.SetFont("dot01")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot02")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot03")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot04")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot05")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot06")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot07")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot08")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot09")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot10")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//     display.Clear()
//     display.SetFont("dot11")
//     display.DrawText(0, 0, "Test")
//     display.Refresh(false)
//     time.Sleep(time.Millisecond * 5000)
//
// 	for i := 0; i < ROW_DRIVER_ROWS; i++ {
// 		display.SetPixel(0, i, true)
// 		display.Refresh(false)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	display.Invert()
// 	display.ShiftHorizontal(5, false)
// 	display.Refresh(false)
// 	time.Sleep(time.Millisecond * 4000)
//
// 	display.ShiftVertical(5, false)
// 	display.Refresh(false)
// 	time.Sleep(time.Millisecond * 4000)
//
// 	display.DrawHorizontalLine(2, 10, 10, true)
// 	display.Refresh(false)
// 	time.Sleep(time.Millisecond * 4000)
//
// 	display.DrawVerticalLine(0, 3, 8, true)
// 	display.Refresh(false)
// 	time.Sleep(time.Millisecond * 4000)
//
// 	display.DrawRectangle(16, 0, 3, 5, true)
// 	display.Refresh(false)
// 	time.Sleep(time.Millisecond * 4000)

//  display.Clear()
// 	display.Refresh(true)
	led.Low()
}
