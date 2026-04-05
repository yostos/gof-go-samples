// Proxy pattern in Go: PrinterProxy lazily initializes a real Printer only
// when Print is called, using sync.Once for thread-safe initialization.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Printable is the subject interface.
type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(s string)
}

// Printer is the real subject that is expensive to create.
type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	p := &Printer{name: name}
	p.heavyJob(fmt.Sprintf("Printerのインスタンス(%s)を生成中", name))
	return p
}

func (p *Printer) SetPrinterName(name string) { p.name = name }
func (p *Printer) GetPrinterName() string      { return p.name }

func (p *Printer) Print(s string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(s)
}

func (p *Printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	fmt.Println("完了。")
}

// PrinterProxy is the proxy that defers Printer creation until Print is called.
type PrinterProxy struct {
	name string
	real *Printer
	mu   sync.Mutex
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

func (pp *PrinterProxy) SetPrinterName(name string) {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	if pp.real != nil {
		pp.real.SetPrinterName(name)
	}
	pp.name = name
}

func (pp *PrinterProxy) GetPrinterName() string {
	return pp.name
}

func (pp *PrinterProxy) Print(s string) {
	pp.realize()
	pp.real.Print(s)
}

func (pp *PrinterProxy) realize() {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	if pp.real == nil {
		pp.real = NewPrinter(pp.name)
	}
}

func main() {
	var p Printable = NewPrinterProxy("Alice")
	fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
	p.SetPrinterName("Bob")
	fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
	p.Print("Hello, world.")
}
