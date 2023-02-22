package main

import (
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

// Start the game and reset global state
func startGame() {
	// Clear display
	clearScreen()
	// Reset Global State
	score = 0
	active = true
	x = 100
	y = 100
	// Display gopher
	tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
}

// Check for collision with compensation
func checkBorder(x, y int16) (int16, int16) {
	if active {
		var border int16 = 10
		// I think compensation is needed due to the size of the gopher
		var x_comp int16 = 40
		var y_comp int16 = 40
		// if hit border, kill
		if (x >= screenWidth-border-x_comp) || (x <= border) || (y <= border+y_comp) || (y >= screenHeight-border) {

			killScreen()
		}
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
