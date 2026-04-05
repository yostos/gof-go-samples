// Decorator pattern in Go: uses a Display interface and wraps it with border
// decorators, similar to how http.Handler middleware works in Go.
package main

import (
	"fmt"
	"strings"
)

// Display is the component interface.
type Display interface {
	Columns() int
	Rows() int
	RowText(row int) string
}

// Show prints all rows of a Display.
func Show(d Display) {
	for i := 0; i < d.Rows(); i++ {
		fmt.Println(d.RowText(i))
	}
}

// StringDisplay is the base concrete component.
type StringDisplay struct {
	str string
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{str: s}
}

func (d *StringDisplay) Columns() int         { return len(d.str) }
func (d *StringDisplay) Rows() int             { return 1 }
func (d *StringDisplay) RowText(row int) string { return d.str }

// SideBorder decorates a Display by adding a character on each side.
type SideBorder struct {
	inner      Display
	borderChar byte
}

func NewSideBorder(d Display, ch byte) *SideBorder {
	return &SideBorder{inner: d, borderChar: ch}
}

func (b *SideBorder) Columns() int { return 1 + b.inner.Columns() + 1 }
func (b *SideBorder) Rows() int    { return b.inner.Rows() }

func (b *SideBorder) RowText(row int) string {
	return string(b.borderChar) + b.inner.RowText(row) + string(b.borderChar)
}

// FullBorder decorates a Display with a full box border.
type FullBorder struct {
	inner Display
}

func NewFullBorder(d Display) *FullBorder {
	return &FullBorder{inner: d}
}

func (b *FullBorder) Columns() int { return 1 + b.inner.Columns() + 1 }
func (b *FullBorder) Rows() int    { return 1 + b.inner.Rows() + 1 }

func (b *FullBorder) RowText(row int) string {
	if row == 0 || row == b.inner.Rows()+1 {
		return "+" + strings.Repeat("-", b.inner.Columns()) + "+"
	}
	return "|" + b.inner.RowText(row-1) + "|"
}

func main() {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1, '#')
	b3 := NewFullBorder(b2)
	Show(b1)
	Show(b2)
	Show(b3)

	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewFullBorder(
						NewStringDisplay("Hello, world."),
					),
					'*',
				),
			),
		),
		'/',
	)
	Show(b4)
}
