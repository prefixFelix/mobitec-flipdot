package matrix

import (
	"errors"
)

const BANDHEIGHT = 5

// Matrix represents a 2D matrix with fixed size, holding values 0, 1, or Empty.
type Matrix struct {
	width, height int
	data          [][]bool
}

//
// Constructors
//

// New creates a new empty Matrix of the given size, initialized to False.
func New(width, height int) *Matrix {
	return &Matrix{width, height, newData(false)}
}

// NewFull creates a new empty Matrix of the given size, initialized to True.
func NewFull(width, height int) *Matrix {
	return &Matrix{width, height, newData(true)}
}

//
// Checked Matrix Manipulation
//

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

//
// Unchecked Matrix Manipulation
//

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

//
// Getter
//

// At returns the value of a cell in the Matrix.
func (m *Matrix) At(row, col int) (bool, error) {
	if row < 0 || row >= Height || col < 0 || col >= Width {
		return false, errors.New("invalid coordinates")
	}
	return m.data[row][col], nil
}

//
// Helper
//

// ToSubcolumn converts the Matrix to a 3D array of subcolumns.
func (m *Matrix) ToSubcolumn() (subcolumnMatrix [][][]bool) {

	// Helper function to create subcolumns for a band
	createBand := func(startRow, endRow int) [][]bool {
		var band [][]bool
		for subcolumns := 0; subcolumns < Width; subcolumns++ {
			var subcolumn []bool
			for row := startRow; row < endRow; row++ {
				val := m.data[row][subcolumns]
				subcolumn = append(subcolumn, val)
			}
			band = append(band, subcolumn)
		}
		return band
	}

	// Process full bands of 5 rows
	for fullBands := 0; fullBands < Height/BANDHEIGHT; fullBands++ {
		startRow := fullBands * BANDHEIGHT
		endRow := startRow + BANDHEIGHT
		band := createBand(startRow, endRow)
		subcolumnMatrix = append(subcolumnMatrix, band)
	}

	// Handle the last band if the height of the matrix is not a multiple of 5
	remainder := Height % BANDHEIGHT
	if remainder != 0 {
		startRow := Height - remainder
		band := createBand(startRow, Height)
		subcolumnMatrix = append(subcolumnMatrix, band)
	}

	return subcolumnMatrix
}

func newData(val bool) [][]bool {
	data := make([][]bool, Height)
	for i := range data {
		data[i] = make([]bool, Width)
		for j := range data[i] {
			data[i][j] = val
		}
	}
	return data
}

func (m *Matrix) getColumn() []bool {
	col := make([]bool, m.height)
	for i := 0; i < m.height; i++ {
		col[i] = false
	}
	return col
}
