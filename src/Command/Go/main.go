// Command pattern: in Go, first-class functions and closures replace Command objects.
// This CLI-based version records draw operations as closures and replays them on a text canvas.
package main

import "fmt"

// Canvas represents a simple text-based drawing surface.
type Canvas struct {
	width, height int
	pixels        [][]bool
}

// NewCanvas creates a canvas of the given size.
func NewCanvas(width, height int) *Canvas {
	pixels := make([][]bool, height)
	for i := range pixels {
		pixels[i] = make([]bool, width)
	}
	return &Canvas{width: width, height: height, pixels: pixels}
}

// Draw sets a pixel at (x, y).
func (c *Canvas) Draw(x, y int) {
	if x >= 0 && x < c.width && y >= 0 && y < c.height {
		c.pixels[y][x] = true
	}
}

// Clear resets all pixels.
func (c *Canvas) Clear() {
	for y := range c.pixels {
		for x := range c.pixels[y] {
			c.pixels[y][x] = false
		}
	}
}

// Print renders the canvas to stdout.
func (c *Canvas) Print() {
	for y := range c.pixels {
		for x := range c.pixels[y] {
			if c.pixels[y][x] {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	canvas := NewCanvas(10, 10)

	// In Go, commands are simply functions (closures).
	// A macro command is just a slice of functions.
	var history []func()

	// Record draw operations (simulating mouse drag drawing a diagonal line)
	points := [][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	for _, p := range points {
		x, y := p[0], p[1]
		cmd := func() { canvas.Draw(x, y) }
		history = append(history, cmd)
		cmd() // execute immediately, like the Java version
		fmt.Printf("Drew dot at (%d, %d)\n", x, y)
	}

	// Record a horizontal line
	for x := 2; x <= 8; x++ {
		xx := x
		cmd := func() { canvas.Draw(xx, 7) }
		history = append(history, cmd)
		cmd()
		fmt.Printf("Drew dot at (%d, %d)\n", xx, 7)
	}

	fmt.Println("\n--- Current Canvas ---")
	canvas.Print()

	// Clear and replay (like the Java repaint that re-executes history)
	fmt.Println("--- Clear and Replay ---")
	canvas.Clear()
	for _, cmd := range history {
		cmd()
	}
	canvas.Print()

	// Undo last 3 commands (pop from history) and replay
	fmt.Println("--- Undo 3 operations and Replay ---")
	history = history[:len(history)-3]
	canvas.Clear()
	for _, cmd := range history {
		cmd()
	}
	canvas.Print()
}
