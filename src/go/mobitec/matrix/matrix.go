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
