package main

import (
	"runtime/interrupt"

	"github.com/scraly/learning-go-by-examples/go-gopher-gba/fonts"
	"tinygo.org/x/tinyfont"
)

func update(interrupt.Interrupt) {

	// check collision
	x, y = checkBorder(x, y)

	// Read uint16 from register regKEYPAD that represents the state of current buttons pressed
	// and compares it against the defined values for each button on the Gameboy Advance
	switch keyValue := regKEYPAD.Get(); keyValue {
	// Start the "game"
	case keySTART:
		startGame()
	// Go back to Menu
	case keySELECT:
		clearScreen()
		drawGophers()
	// Gopher go to the right
	case keyRIGHT:
		x, y = move(x, y, 10, false, true)
	// Gopher go to the left
	case keyLEFT:
		x, y = move(x, y, 10, false, false)
	// Gopher go to the down
	case keyDOWN:
		x, y = move(x, y, 10, true, true)
	case keyUP:
		x, y = move(x, y, 10, true, false)
	//Gopher jump
	case keyA:
		x, y = move(x, y, 20, true, false)
		// Clear the display
		x, y = move(x, y, 20, true, true)
	}
	// Add random movement
	x, y = wind(x, y)

	// Increment Global Counter
	score = score + 1

}

// Add random movement to bottom left of screen
func wind(x, y int16) (int16, int16) {
	if active {
		// increase wind_power as score goes higher
		var wind_power int16 = score / 10

		x, y = move(x, y, wind_power, random(), random())
	}
	return x, y
}

// importing math/random overflew my game
func random() bool {
	if score%2 == 0 {
		return false
	} else {
		return true
	}
}

// move abstraction
func move(current_x, current_y, pixels int16, vertical, positive bool) (int16, int16) {
	// Clear display by drawing a gopher in black
	tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)

	if vertical {
		if positive {
			y = y + pixels
		} else {
			y = y - pixels
		}
	} else {
		if positive {
			x = x + pixels
		} else {
			x = x - pixels
		}
	}
	tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	return x, y
}
