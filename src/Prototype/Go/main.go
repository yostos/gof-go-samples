// Prototype: Goでは値セマンティクスにより構造体のコピーは代入で完了するため、
// Cloneableインターフェースやプロトタイプレジストリは不要。構造体をそのままコピーして使う。
package main

import (
	"fmt"
	"strings"
)

// Decorator defines the interface for text decorations.
type Decorator interface {
	Use(s string)
}

// MessageBox decorates text inside a box of characters.
type MessageBox struct {
	DecoChar rune
}

func (m MessageBox) Use(s string) {
	decoLen := 1 + len(s) + 1
	line := strings.Repeat(string(m.DecoChar), decoLen)
	fmt.Println(line)
	fmt.Printf("%c%s%c\n", m.DecoChar, s, m.DecoChar)
	fmt.Println(line)
}

// UnderlinePen decorates text with an underline.
type UnderlinePen struct {
	ULChar rune
}

func (u UnderlinePen) Use(s string) {
	fmt.Println(s)
	fmt.Println(strings.Repeat(string(u.ULChar), len(s)))
}

func main() {
	// In Go, value types are copied by simple assignment.
	// No clone/prototype registry needed -- just copy the struct.
	upen := UnderlinePen{ULChar: '-'}
	mbox := MessageBox{DecoChar: '*'}
	sbox := MessageBox{DecoChar: '/'}

	// Register prototypes in a map (same concept as Java's Manager).
	// The map stores copies of the values.
	prototypes := map[string]Decorator{
		"strong message": upen,
		"warning box":    mbox,
		"slash box":      sbox,
	}

	// Create copies by reading from the map.
	// In Go, map lookup of a value type returns a copy automatically.
	p1 := prototypes["strong message"]
	p1.Use("Hello, world.")

	p2 := prototypes["warning box"]
	p2.Use("Hello, world.")

	p3 := prototypes["slash box"]
	p3.Use("Hello, world.")
}
