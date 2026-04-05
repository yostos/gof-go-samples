// Adapter pattern in Go: wraps a Banner struct to satisfy the Printer interface
// using struct embedding, which is the idiomatic Go way to adapt interfaces.
package main

import "fmt"

// Printer is the target interface that clients expect.
type Printer interface {
	PrintWeak()
	PrintStrong()
}

// Banner is the existing (adaptee) type with its own method names.
type Banner struct {
	str string
}

func NewBanner(s string) *Banner {
	return &Banner{str: s}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.str)
}

// BannerAdapter adapts Banner to the Printer interface by embedding it
// and delegating to its methods.
type BannerAdapter struct {
	*Banner
}

func NewBannerAdapter(s string) *BannerAdapter {
	return &BannerAdapter{Banner: NewBanner(s)}
}

func (a *BannerAdapter) PrintWeak() {
	a.ShowWithParen()
}

func (a *BannerAdapter) PrintStrong() {
	a.ShowWithAster()
}

func main() {
	var p Printer = NewBannerAdapter("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
