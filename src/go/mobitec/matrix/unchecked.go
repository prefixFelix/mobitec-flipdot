package matrix

var currentValue = true

//
// General
//

// Fill fills the Matrix.
func (m *Matrix) Fill() {
	for i := range m.data {
		for j := range m.data[i] {
			m.data[i][j] = true
		}
	}
}

// Clear clears the Matrix.
func (m *Matrix) Clear() {
	for i := range m.data {
		for j := range m.data[i] {
			m.data[i][j] = false
		}
	}
}

type Offset struct {
	X  int
	Y  int
	Re int
}

// DrawPath draws a path in the Matrix based on a start position and a list of offsets.
// Offsets are a list of Os, which are a struct with X and Y values and a Re value.
// The Re value indicates how many times the offset should be applied.
// A pixel is drawn at every step.
//
// For example:
// DrawPath(0, 0, []matrix.Offset{{X: 1}, {Y: -1}, {X: 1, Re: 4}, {X: 1, Y: 1, Re: 3}})
// Would draw one pixel to the right, one up, four to the right, then diagonally down-right three pixels.
func (m *Matrix) DrawPath(startX, startY int, offsets []Offset) {
	m.Set(startX, startY)

	// Apply each offset to the current position and set the value
	for _, offset := range offsets {
		if offset.Re == 0 {
			offset.Re = 1
		}
		for i := 0; i < offset.Re; i++ {
			startX += offset.X
			startY += offset.Y
			m.Set(startX, startY)
		}
	}
}

//
// Setters
//

// SetValue sets which value is to be used for the manipulation functions. It's like choosing the color of your pen.
func (m *Matrix) SetValue(value bool) {
	currentValue = value
}

// Set sets the value of a pixel in the Matrix.
func (m *Matrix) Set(col, row int) {
	m.data[row][col] = currentValue
}

// SetColumn sets the value of a column in the Matrix.
func (m *Matrix) SetColumn(colIdx int) {
	for rowIdx := range m.data {
		m.data[rowIdx][colIdx] = currentValue
	}
}

// SetRow sets the value of a row in the Matrix.
func (m *Matrix) SetRow(rowIdx int) {
	for colIdx := range m.data[rowIdx] {
		m.data[rowIdx][colIdx] = currentValue
	}
}

// SetColumnFromTo sets the value of the given column from a given row to another.
// To can be -1 to indicate the end of the column.
func (m *Matrix) SetColumnFromTo(colIdx, from, to int) {
	if to == -1 {
		to = m.Height
	}
	for rowIdx := from; rowIdx < to; rowIdx++ {
		m.data[rowIdx][colIdx] = currentValue
	}
}

// SetRowFromTo sets the value of the given row from a given column to another.
// To can be -1 to indicate the end of the row.
func (m *Matrix) SetRowFromTo(rowIdx, from, to int) {
	if to == -1 {
		to = m.Width
	}
	for colIdx := from; colIdx < to; colIdx++ {
		m.data[rowIdx][colIdx] = currentValue
	}
}

// FillBounds fills all pixels connected to the (x, y) pixel until it meets boundaries.
// Boundaries meaning either the edge of the display or other connected pixels
func (m *Matrix) FillBounds(x, y int) {
	m.fillHelper(x, y)
}

//
// Other Manipulation
//

// PutColumn sets the value of a column in the Matrix.
func (m *Matrix) PutColumn(colIdx int, col []bool) {
	for rowIdx, val := range col {
		m.data[rowIdx][colIdx] = val
	}
}

// PutRow sets the value of a row in the Matrix.
func (m *Matrix) PutRow(rowIdx int, row []bool) {
	copy(m.data[rowIdx], row)
}

// RepeatForColumn sets the value of a column in the Matrix to a repeating pattern with a variable length.
func (m *Matrix) RepeatForColumn(colIdx int, values []bool) {
	i := 0
	for rowIdx := 0; rowIdx < len(m.data); rowIdx++ {
		if i == len(values) {
			i = 0
		}
		m.data[rowIdx][colIdx] = values[i]
		i++
	}
}

// RepeatForRow sets the value of a row in the Matrix to a repeating pattern with a variable length.
func (m *Matrix) RepeatForRow(rowIdx int, values []bool) {
	i := 0
	for colIdx := 0; colIdx < len(m.data[rowIdx]); colIdx++ {
		if i == len(values) {
			i = 0
		}
		m.data[rowIdx][colIdx] = values[i]
	}
}

// ShiftLeft shifts the Matrix to the left by one column.
func (m *Matrix) ShiftLeft() {
	for i := range m.data {
		m.data[i] = m.data[i][1:]
		m.data[i] = append(m.data[i], m.getColumn()...)
	}
}

// ShiftRight shifts the Matrix to the right by one column.
func (m *Matrix) ShiftRight() {
	for i := range m.data {
		m.data[i] = append(m.getColumn(), m.data[i][:m.Width-1]...)
	}
}

// Negate negates the Matrix.
func (m *Matrix) Negate() {
	for i := range m.data {
		for j := range m.data[i] {
			m.data[i][j] = !m.data[i][j]
		}
	}
}

//
// Helper
//

// fillHelper is a recursive helper function to implement the fill functionality.
func (m *Matrix) fillHelper(x, y int) {
	// Base case: Check if pixel is out of bounds or already filled/ a wall
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height || m.data[y][x] {
		return
	}

	// Fill the current pixel
	m.data[y][x] = currentValue

	// Recursively fill the neighboring pixels
	m.fillHelper(x+1, y) // Right
	m.fillHelper(x-1, y) // Left
	m.fillHelper(x, y+1) // Down
	m.fillHelper(x, y-1) // Up
}
