package main

import (
	"runtime/interrupt"
	"strconv"

	"github.com/scraly/learning-go-by-examples/go-gopher-gba/fonts"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

func drawGophers() {
	// Display a textual message "Gopher" with Google colors
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 36, 60, 'G', gBlue)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 71, 60, 'o', gRed)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 98, 60, 'p', gYellow)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 126, 60, 'h', gGreen)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 154, 60, 'e', gBlue)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 180, 60, 'r', gRed)

	// Display a "press START button" message - center
	tinyfont.WriteLine(display, &tinyfont.TomThumb, 85, 90, "Press START button", white)

	// Display a textual message "Zuri" with Google colors
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 71, 130, 'Z', gBlue)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 98, 130, 'u', gRed)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 126, 130, 'r', gYellow)
	tinyfont.DrawChar(display, &fonts.Bold24pt7b, 144, 130, 'i', gGreen)

	// Display two gophers
	tinyfont.DrawChar(display, &fonts.Regular58pt, 5, 140, 'J', green)
	tinyfont.DrawChar(display, &fonts.Regular58pt, 195, 140, 'X', red)

	tinydraw.Rectangle(display, int16(0), int16(0), screenWidth, screenHeight, red)

}

func update(interrupt.Interrupt) {

	// Read uint16 from register regKEYPAD that represents the state of current buttons pressed
	// and compares it against the defined values for each button on the Gameboy Advance
	switch keyValue := regKEYPAD.Get(); keyValue {
	// Start the "game"
	case keySTART:
		// Clear display
		clearScreen()
		// Reset Global State
		score = 0
		active = true
		// Display gopher
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
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

	// check collision
	x, y = checkBorder(x, y)

	// Increment Global Counter
	score = score + 1

}

// Add random movement to bottom left of screen
func wind(x, y int16) (int16, int16) {
	if active {
		// increase wind_power as score goes higher
		var wind_power int16 = int16(score / 20)

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

// Overwrite the entire screen
func clearScreen() {
	tinydraw.FilledRectangle(
		display,
		int16(0), int16(0),
		screenWidth, screenHeight,
		black,
	)
	tinydraw.Rectangle(display, int16(0), int16(0), screenWidth, screenHeight, red)

}

// Check for collision with compensation
func checkBorder(x, y int16) (int16, int16) {
	var border int16 = 10
	// I think compensation is needed due to the size of the gopher
	var x_comp int16 = 40
	var y_comp int16 = 40
	// if hit border, kill
	if (x >= screenWidth-border-x_comp) || (x <= border) || (y <= border+y_comp) || (y >= screenHeight-border) {
		x = 100
		y = 100
		killScreen()
	}
	return x, y

}

// Display end screen and pause game
func killScreen() {
	for {
		clearScreen()
		tinyfont.WriteLine(display, &tinyfont.TomThumb, 105, 30, "You DIED!", red)
		tinyfont.WriteLine(display, &tinyfont.TomThumb, 110, 45, "Score", red)
		tinyfont.WriteLine(display, &tinyfont.TomThumb, 115, 55, strconv.Itoa(int(score)), red)
		tinyfont.WriteLine(display, &tinyfont.TomThumb, 80, 130, "Press start to restart", red)
		tinyfont.DrawChar(display, &fonts.Regular58pt, 95, 110, 'C', red)
		active = false
		if regKEYPAD.Get() == keySELECT {
			break
		}
		break
		clearScreen()

	}
}
