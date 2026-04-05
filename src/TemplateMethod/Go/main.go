// Template Method pattern is unnecessary in Go since Go has no abstract classes.
// Instead, we use an interface for the varying steps and a plain function for the template.
package main

import (
	"fmt"
	"strings"
)

// Displayer defines the varying steps that each display type implements.
type Displayer interface {
	Open()
	Print()
	Close()
}

// Display is the template function that calls open, print (5 times), and close.
func Display(d Displayer) {
	d.Open()
	for i := 0; i < 5; i++ {
		d.Print()
	}
	d.Close()
}

// CharDisplay displays a single character with << >> brackets.
type CharDisplay struct {
	ch byte
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}

func (c *CharDisplay) Print() {
	fmt.Printf("%c", c.ch)
}

func (c *CharDisplay) Close() {
	fmt.Println(">>")
}

// StringDisplay displays a string in a box made of +---+ and | |.
type StringDisplay struct {
	str   string
	width int
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{str: s, width: len(s)}
}

func (sd *StringDisplay) printLine() {
	fmt.Printf("+%s+\n", strings.Repeat("-", sd.width))
}

func (sd *StringDisplay) Open() {
	sd.printLine()
}

func (sd *StringDisplay) Print() {
	fmt.Printf("|%s|\n", sd.str)
}

func (sd *StringDisplay) Close() {
	sd.printLine()
}

func main() {
	d1 := &CharDisplay{ch: 'H'}
	d2 := NewStringDisplay("Hello, world.")

	Display(d1)
	Display(d2)
}
