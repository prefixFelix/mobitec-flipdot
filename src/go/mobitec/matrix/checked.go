package matrix

import "errors"

// SetChecked sets the value of a pixel in the Matrix, if the coordinates and value are valid.
func (m *Matrix) SetChecked(row, col int, value bool) error {
	if row < 0 || row >= m.Height || col < 0 || col >= m.Width {
		return errors.New("invalid coordinates")
	}
	m.Set(row, col)
	return nil
}

// PutColumnChecked sets the value of a column in the Matrix, if the index and column are valid.
func (m *Matrix) PutColumnChecked(colIdx int, col []bool) error {
	if colIdx < 0 || colIdx >= m.Width {
		return errors.New("invalid column index")
	}
	if len(col) != m.Height {
		return errors.New("invalid column length")
	}

	m.PutColumn(colIdx, col)
	return nil
}

// PutRowChecked sets the value of a row in the Matrix, if the index and row are valid.
func (m *Matrix) PutRowChecked(rowIdx int, row []bool) error {
	if rowIdx < 0 || rowIdx >= m.Height {
		return errors.New("invalid row index")
	}
	if len(row) != m.Width {
		return errors.New("invalid row length")
	}

	m.PutRow(rowIdx, row)
	return nil
}

// SetColumnChecked sets the value of a column in the Matrix, if the index is valid.
func (m *Matrix) SetColumnChecked(colIdx int, value bool) error {
	if colIdx < 0 || colIdx >= m.Width {
		return errors.New("invalid column index")
	}

	m.SetColumn(colIdx)
	return nil
}

// SetRowChecked sets the value of a row in the Matrix, if the index is valid.
func (m *Matrix) SetRowChecked(rowIdx int) error {
	if rowIdx < 0 || rowIdx >= m.Height {
		return errors.New("invalid row index")
	}

	m.SetRow(rowIdx)
	return nil
}

// SetRowFromToChecked sets the value of the given row from a given column to another, if the indexes are valid.
// To can be -1 to indicate the end of the row.
func (m *Matrix) SetRowFromToChecked(rowIdx, from, to int) error {
	if rowIdx < 0 || rowIdx >= m.Height || from < 0 || from >= m.Width || to < -1 || to >= m.Width {
		return errors.New("invalid indexes")
	}

	m.SetRowFromTo(rowIdx, from, to)
	return nil
}

// SetColumnFromToChecked sets the value of the given column from a given row to another, if the indexes are valid.
// To can be -1 to indicate the end of the column.
func (m *Matrix) SetColumnFromToChecked(colIdx, from, to int) error {
	if colIdx < 0 || colIdx >= m.Width || from < 0 || from >= m.Height || to < -1 || to >= m.Height {
		return errors.New("invalid indexes")
	}

	m.SetColumnFromTo(colIdx, from, to)
	return nil
}

// RepeatForColumnChecked sets the value of a column in the Matrix to a repeating pattern with a variable length, if the index and values are valid.
func (m *Matrix) RepeatForColumnChecked(colIdx int, values []bool) error {
	if colIdx < 0 || colIdx >= m.Width {
		return errors.New("invalid column index")
	}
	if len(values) != m.Height {
		return errors.New("invalid values length")
	}

	m.RepeatForColumn(colIdx, values)
	return nil
}

// RepeatForRowChecked sets the value of a row in the Matrix to a repeating pattern with a variable length, if the index and values are valid.
func (m *Matrix) RepeatForRowChecked(rowIdx int, values []bool) error {
	if rowIdx < 0 || rowIdx >= m.Height {
		return errors.New("invalid row index")
	}
	if len(values) != m.Width {
		return errors.New("invalid values length")
	}

	m.RepeatForRow(rowIdx, values)
	return nil
}

// FillBoundsChecked fills all pixels connected to the (x, y) pixel until it meets boundaries.
func (m *Matrix) FillBoundsChecked(x, y int) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("coordinates out of bounds")
	}
	m.Fill()
	return nil
}
