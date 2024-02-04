package matrix

// SetChecked the display dimensions.
// Happens automatically when mobitec.Display is initialized.
var (
	Height int
	Width  int
)

// NewDisplayMatrix creates a new empty display matrix.
// A DisplayMatrix is a matrix that automatically takes on the dimensions of the display.
func NewDisplayMatrix() *Matrix {
	return New(Width, Height)
}

func NewFullDisplayMatrix() *Matrix {
	return NewFull(Width, Height)
}
