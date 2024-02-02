package matrix

import (
	"errors"
)

var (
	Height int
	Width  int
)

const BANDHEIGHT = 5

// Matrix represents a 2D matrix with fixed size, holding values 0, 1, or Empty.
type Matrix struct {
	data [][]bool
}

// New creates a new empty Matrix of the given size, initialized to False.
func New() *Matrix {
	return &Matrix{newData(false)}
}

// NewFull creates a new empty Matrix of the given size, initialized to True.
func NewFull() *Matrix {
	return &Matrix{newData(true)}
}

// Set sets the value of a cell in the Matrix, if the coordinates and value are valid.
func (m *Matrix) Set(row, col int, value bool) error {
	if row < 0 || row >= Height || col < 0 || col >= Width {
		return errors.New("invalid coordinates")
	}

	m.data[row][col] = value
	return nil
}

// Get returns the value of a cell in the Matrix.
func (m *Matrix) Get(row, col int) (bool, error) {
	if row < 0 || row >= Height || col < 0 || col >= Width {
		return false, errors.New("invalid coordinates")
	}
	return m.data[row][col], nil
}

func (m *Matrix) SetColumn(colIdx int, col []bool) error {
	if colIdx < 0 || colIdx >= Width {
		return errors.New("invalid column index")
	}
	if len(col) != Height {
		return errors.New("invalid column length")
	}

	for rowIdx, val := range col {
		m.data[rowIdx][colIdx] = val
	}
	return nil
}

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
