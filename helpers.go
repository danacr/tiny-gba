package main

import (
	"runtime/interrupt"
	"time"

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
	tinyfont.DrawChar(display, &fonts.Regular58pt, 5, 140, 'B', green)
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
		// Display gopher
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Go back to Menu
	case keySELECT:
		clearScreen()
		drawGophers()
	// Gopher go to the right
	case keyRIGHT:
		// Clear display by drawing a gopher in black
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		x = x + 10
		// display gopher at right
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the left
	case keyLEFT:
		// Clear display
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		x = x - 10
		// display gopher at right
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the down
	case keyDOWN:
		// Clear display
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		y = y + 10
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the up
	case keyUP:
		// Clear display
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		y = y - 10
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	//Gopher jump
	case keyA:
		// Clear display
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		// Display the gopher up
		y = y - 20
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
		// Clear the display
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', black)
		// Display the gopher down
		y = y + 20
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	}
	x, y = checkBorder(x, y)

}

func clearScreen() {
	tinydraw.FilledRectangle(
		display,
		int16(0), int16(0),
		screenWidth, screenHeight,
		black,
	)
	tinydraw.Rectangle(display, int16(0), int16(0), screenWidth, screenHeight, red)

}

func checkBorder(x, y int16) (int16, int16) {
	var border int16 = 10
	// I think compensation is needed due to the size of the gopher
	var x_comp int16 = 40
	var y_comp int16 = 40
	// if hit border, kill
	if (x >= screenWidth-border-x_comp) || (x <= border) || (y <= border+y_comp) || (y >= screenHeight-border) {
		clearScreen()
		tinyfont.WriteLine(display, &tinyfont.TomThumb, 85, 90, "You DIED!", red)
		x = 100
		y = 100
		time.Sleep(time.Second * 3)
		clearScreen()
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	}
	return x, y

}
