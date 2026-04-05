// Bridge pattern in Go: since Go has no class hierarchy, we simply use interfaces
// and composition. The "bridge" is just normal Go design with an interface dependency.
package main

import (
	"fmt"
	"strings"
)

// DisplayImpl is the implementor interface.
type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

// Display uses a DisplayImpl via composition (the "bridge" in Java terms).
func display(impl DisplayImpl) {
	impl.RawOpen()
	impl.RawPrint()
	impl.RawClose()
}

// multiDisplay prints the content multiple times.
func multiDisplay(impl DisplayImpl, times int) {
	impl.RawOpen()
	for i := 0; i < times; i++ {
		impl.RawPrint()
	}
	impl.RawClose()
}

// StringDisplayImpl is a concrete implementor that draws a bordered string.
type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(s string) *StringDisplayImpl {
	return &StringDisplayImpl{str: s, width: len(s)}
}

func (d *StringDisplayImpl) RawOpen() {
	d.printLine()
}

func (d *StringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", d.str)
}

func (d *StringDisplayImpl) RawClose() {
	d.printLine()
}

func (d *StringDisplayImpl) printLine() {
	fmt.Printf("+%s+\n", strings.Repeat("-", d.width))
}

func main() {
	d1 := NewStringDisplayImpl("Hello, Japan.")
	d2 := NewStringDisplayImpl("Hello, World.")
	d3 := NewStringDisplayImpl("Hello, Universe.")

	display(d1)
	display(d2)
	display(d3)
	multiDisplay(d3, 5)
}
