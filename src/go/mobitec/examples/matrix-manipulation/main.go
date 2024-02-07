package main

import (
	"fmt"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec"
	. "github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/examples"
	"github.com/prefixFelix/mobitec-flipdot/src/go/mobitec/matrix"
)

var display *mobitec.Display

func main() {
	// Set the display values in cfg.go
	//
	// Initialize the display, open the serial connection and make sure it is closed at the end
	display = mobitec.NewDisplay(PORT, WIDTH, HEIGHT)
	err := display.Open()
	Check(err)
	defer display.Close()

	// A DisplayMatrix is a Matrix that automatically gets the correct size for the display
	m := matrix.NewDisplayMatrix()

	// "warm up" the display by setting/resetting all dots
	m.Fill()

	// Use display.Matrix to send the matrix to the display
	err = display.Matrix(*m)
	Check(err)

	m.Clear()
	toDisplay(m)

	// Put a column of pixels into the matrix
	// SetValue is like choosing the color of your pen. It's the value most manipulation functions will use.
	// True means dots will be set, false means they will be cleared. It's true by default.
	m.SetValue(true)
	m.SetColumn(WIDTH / 2)

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
	m.Set(0, 0)
	toDisplay(m)

	// Set a row/column to a slice of values
	m.PutColumn(
		1,
		[]bool{true, true, false, true, false, true, true, false, false, false, false, false, false, true, true, true},
	)
	toDisplay(m)

	// Set a full row/column to a value
	m.SetRow(HEIGHT / 2)
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
	m.SetValue(false)
	for i := drawFrom - 1; i <= drawTo+1; i++ {
		m.SetColumn(i)
	}
	toDisplay(m)

	// Draw a simple shape
	m.SetValue(true)
	for i := drawFrom; i < WIDTH/2; i++ {
		m.SetColumnFromTo(i, i-drawFrom, HEIGHT-i+drawFrom)
	}
	for i := drawTo; i >= WIDTH/2; i-- {
		m.SetColumnFromTo(i, (i-drawTo)*-1, HEIGHT+i-drawTo)
	}
	toDisplay(m)
	Wait()
	Wait()

	// Draw a rocket ship
	ROCKETSHIP_LENGTH := WIDTH
	m.Clear()

	// Lower Part
	m.DrawPath(0, 1, []matrix.Offset{
		{X: 1},
		{X: 1, Y: 1},
		{X: 1},
		{X: 1, Y: 1, Re: 2},
		{Y: 1},
		{X: -1, Y: 1, Re: 2},
		{X: -1},
		{X: -2, Y: 1},
		{X: 1, Re: 3},
		{X: -1, Y: 1},
		{X: 1},
		{X: 1, Y: 1, Re: 2},
		{Y: 1},
		{X: -1, Y: 1, Re: 2},
		{X: -1},
		{X: -1, Y: 1},
		{X: -1},
	})

	end := ROCKETSHIP_LENGTH - 17
	// Top: Bottom Part
	m.SetRowFromTo(2, 4, end+5)
	m.SetRowFromTo(14, 4, end+5)
	m.DrawPath(end+4, 3, []matrix.Offset{
		{X: -1, Y: 1},
		{Y: 1, Re: 8},
		{X: 1, Y: 1},
	})

	// Top: Top Part
	m.DrawPath(end+5, 3, []matrix.Offset{
		{X: 1, Re: 4},
		{X: 1, Y: 1},
		{X: 1},
		{X: 1, Y: 1, Re: 3},
		{X: -5, Y: 1},
		{X: 1, Re: 5},
		{Y: 1},
		{X: -1, Y: 1, Re: 3},
		{X: -1},
		{X: -1, Y: 1},
		{X: -1, Re: 4},
	})

	toDisplay(m)
}

func toDisplay(m *matrix.Matrix) {
	Check(display.Matrix(*m))
	Wait()
}
