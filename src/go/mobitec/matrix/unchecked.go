package matrix

// Set sets the value of a pixel in the Matrix.
func (m *Matrix) Set(row, col int, value bool) {
	m.data[row][col] = value
}

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

// SetColumn sets the value of a column in the Matrix.
func (m *Matrix) SetColumn(colIdx int, value bool) {
	for rowIdx := range m.data {
		m.data[rowIdx][colIdx] = value
	}
}

// SetRow sets the value of a row in the Matrix.
func (m *Matrix) SetRow(rowIdx int, value bool) {
	for colIdx := range m.data[rowIdx] {
		m.data[rowIdx][colIdx] = value
	}
}

// SetColumnFromTo sets the value of the given column from a given row to another.
func (m *Matrix) SetColumnFromTo(colIdx, from, to int, value bool) {
	for rowIdx := from; rowIdx < to; rowIdx++ {
		m.data[rowIdx][colIdx] = value
	}
}

// SetRowFromTo sets the value of the given row from a given column to another.
func (m *Matrix) SetRowFromTo(rowIdx, from, to int, value bool) {
	for colIdx := from; colIdx < to; colIdx++ {
		m.data[rowIdx][colIdx] = value
	}
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
		m.data[i] = append(m.getColumn(), m.data[i][:Width-1]...)
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
