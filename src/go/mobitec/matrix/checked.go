package matrix

import "errors"

// SetChecked sets the value of a pixel in the Matrix, if the coordinates and value are valid.
func (m *Matrix) SetChecked(row, col int, value bool) error {
	if row < 0 || row >= Height || col < 0 || col >= Width {
		return errors.New("invalid coordinates")
	}
	m.Set(row, col, value)
	return nil
}

// PutColumnChecked sets the value of a column in the Matrix, if the index and column are valid.
func (m *Matrix) PutColumnChecked(colIdx int, col []bool) error {
	if colIdx < 0 || colIdx >= Width {
		return errors.New("invalid column index")
	}
	if len(col) != Height {
		return errors.New("invalid column length")
	}

	m.PutColumn(colIdx, col)
	return nil
}

// PutRowChecked sets the value of a row in the Matrix, if the index and row are valid.
func (m *Matrix) PutRowChecked(rowIdx int, row []bool) error {
	if rowIdx < 0 || rowIdx >= Height {
		return errors.New("invalid row index")
	}
	if len(row) != Width {
		return errors.New("invalid row length")
	}

	m.PutRow(rowIdx, row)
	return nil
}

// SetColumnChecked sets the value of a column in the Matrix, if the index is valid.
func (m *Matrix) SetColumnChecked(colIdx int, value bool) error {
	if colIdx < 0 || colIdx >= Width {
		return errors.New("invalid column index")
	}

	m.SetColumn(colIdx, value)
	return nil
}

// SetRowChecked sets the value of a row in the Matrix, if the index is valid.
func (m *Matrix) SetRowChecked(rowIdx int, value bool) error {
	if rowIdx < 0 || rowIdx >= Height {
		return errors.New("invalid row index")
	}

	m.SetRow(rowIdx, value)
	return nil
}

// SetRowFromToChecked sets the value of the given row from a given column to another, if the indexes are valid.
func (m *Matrix) SetRowFromToChecked(rowIdx, from, to int, value bool) error {
	if rowIdx < 0 || rowIdx >= Height || from < 0 || from >= Width || to < 0 || to >= Width {
		return errors.New("invalid indexes")
	}

	m.SetRowFromTo(rowIdx, from, to, value)
	return nil
}

// SetColumnFromToChecked sets the value of the given column from a given row to another, if the indexes are valid.
func (m *Matrix) SetColumnFromToChecked(colIdx, from, to int, value bool) error {
	if colIdx < 0 || colIdx >= Width || from < 0 || from >= Height || to < 0 || to >= Height {
		return errors.New("invalid indexes")
	}

	m.SetColumnFromTo(colIdx, from, to, value)
	return nil
}

// RepeatForColumnChecked sets the value of a column in the Matrix to a repeating pattern with a variable length, if the index and values are valid.
func (m *Matrix) RepeatForColumnChecked(colIdx int, values []bool) error {
	if colIdx < 0 || colIdx >= Width {
		return errors.New("invalid column index")
	}
	if len(values) != Height {
		return errors.New("invalid values length")
	}

	m.RepeatForColumn(colIdx, values)
	return nil
}

// RepeatForRowChecked sets the value of a row in the Matrix to a repeating pattern with a variable length, if the index and values are valid.
func (m *Matrix) RepeatForRowChecked(rowIdx int, values []bool) error {
	if rowIdx < 0 || rowIdx >= Height {
		return errors.New("invalid row index")
	}
	if len(values) != Width {
		return errors.New("invalid values length")
	}

	m.RepeatForRow(rowIdx, values)
	return nil
}
