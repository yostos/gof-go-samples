// Singleton pattern in Go: instead of a global getInstance(), the single instance
// is created once and shared via dependency injection (DI).
// This is the modern Go idiom — no hidden global state, explicit dependencies.
package main

import "fmt"

// Singleton corresponds to the Java Singleton class.
// In Go, it is a regular exported struct; the "single instance" guarantee is
// achieved by creating it once in main and passing it where needed, not by
// hiding the constructor.
type Singleton struct{}

func NewSingleton() *Singleton {
	fmt.Println("インスタンスを生成しました。")
	return &Singleton{}
}

func main() {
	fmt.Println("Start.")

	// Create one instance and inject it wherever needed.
	// In Go, the single-instance guarantee comes from discipline at the
	// composition root (main), not from a private constructor.
	obj1 := NewSingleton()
	obj2 := obj1 // DI: share the same instance by passing the pointer

	if obj1 == obj2 {
		fmt.Println("obj1とobj2は同じインスタンスです。")
	} else {
		fmt.Println("obj1とobj2は同じインスタンスではありません。")
	}

	fmt.Println("End.")
}
