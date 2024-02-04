package fonts

import (
	"github.com/ByteSizedMarius/bdf-explorer/bdf"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/matrix"
)

type CustomFont struct {
	Font *bdf.Font
}

func NewCustom(path string) (cf CustomFont, err error) {
	font, err := bdf.FromFile(path)
	if err != nil {
		return
	}

	return CustomFont{Font: font}, nil
}

func (cf CustomFont) GetMatrix(text string) *matrix.Matrix {
	// Initialize an empty matrix
	m := matrix.NewDisplayMatrix()

	// Initialize the current position
	curX := 0

	// Iterate over each character in the text
	for _, char := range text {
		// At the corresponding glyph from the BDF font
		glyph, exists := cf.Font.CharMap[char]
		if !exists {
			continue
		}

		// Add the glyph to the matrix bitwise
		bounds := glyph.Alpha.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				_, _, _, a := glyph.Alpha.At(x, y).RGBA()
				if a > 0 {
					err := m.SetChecked(y, curX+x, true)
					if err != nil {
						continue
					}
				}
			}
		}

		// Move the current position down by the height of the glyph minus the offset
		curX += bounds.Dx()
	}

	return m
}
