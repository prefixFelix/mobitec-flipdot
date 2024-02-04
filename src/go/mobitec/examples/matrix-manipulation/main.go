package main

import (
	"fmt"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec"
	. "github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/examples"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/matrix"
	"time"
)

var display *mobitec.Display

func wait() {
	time.Sleep(750 * WIDTH / 28 * time.Millisecond)

}

func toDisplay(m *matrix.Matrix) {
	Check(display.Matrix(*m))
	wait()
}

func main() {
	display = mobitec.NewDisplay(PORT, WIDTH, HEIGHT)
	err := display.Open()
	Check(err)

	// "warm up" the display
	err = display.Fill()
	Check(err)
	wait()
	err = display.Clear()
	Check(err)
	wait()

	// A DisplayMatrix is a Matrix that automatically gets the correct size for the display
	m := matrix.NewDisplayMatrix()

	// Put a column of pixels into the matrix
	m.SetColumn(WIDTH/2, true)

	// Set the matrix into the display (send the matrix to the display)
	// All pixels that are not set will be cleared
	err = display.Matrix(*m)
	Check(err)

	// If your code is uncertain about the bounds of the display, you can use checked operations
	// This will avoid panics
	err = m.SetColumnChecked(29, true)
	if err != nil {
		fmt.Println("You seem to have small one...", err.Error())
	} else {
		toDisplay(m)
	}

	//
	// Matrix Manipulation
	//

	// Set a single pixel
	m.Set(0, 0, true)
	toDisplay(m)

	// Set a row/column to a slice of values
	m.PutColumn(
		1,
		[]bool{true, true, false, true, false, true, true, false, false, false, false, false, false, true, true, true},
	)
	toDisplay(m)

	// Set a full row/column to a value
	m.SetRow(HEIGHT/2, true)
	toDisplay(m)

	// Set a repeating pattern
	m.RepeatForColumn(3, []bool{true, false})
	toDisplay(m)

	// Shift left/right (no loop around)
	m.ShiftLeft()
	toDisplay(m)

	// Negate
	m.Negate()
	toDisplay(m)

	// Clear a part of the display
	drawWidth := 14
	drawFrom := WIDTH/2 - drawWidth/2
	drawTo := drawFrom + drawWidth
	for i := drawFrom - 1; i <= drawTo+1; i++ {
		m.SetColumn(i, false)
	}
	toDisplay(m)

	// Draw a simple shape
	for i := drawFrom; i < WIDTH/2; i++ {
		m.SetColumnFromTo(i, i-drawFrom, HEIGHT-i+drawFrom, true)
	}
	for i := drawTo; i >= WIDTH/2; i-- {
		m.SetColumnFromTo(i, (i-drawTo)*-1, HEIGHT+i-drawTo, true)
	}

	toDisplay(m)
}
