// Chain of Responsibility pattern in Go: a linked list of Support handlers.
// Each handler tries to resolve a Trouble; if it cannot, it passes to the next.
package main

import "fmt"

// Trouble represents a problem identified by a number.
type Trouble struct {
	Number int
}

func (t Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.Number)
}

// Support is the handler interface in the chain.
type Support interface {
	SetNext(next Support) Support
	Handle(trouble Trouble)
	Resolve(trouble Trouble) bool
	String() string
}

// BaseSupport provides the common chain logic (name, next pointer, Handle).
// Concrete handlers embed this and implement Resolve.
type BaseSupport struct {
	name    string
	next    Support
	resolve func(Trouble) bool
}

func (s *BaseSupport) SetNext(next Support) Support {
	s.next = next
	return next
}

func (s *BaseSupport) Handle(trouble Trouble) {
	if s.resolve(trouble) {
		fmt.Printf("%s is resolved by [%s].\n", trouble, s.name)
	} else if s.next != nil {
		s.next.Handle(trouble)
	} else {
		fmt.Printf("%s cannot be resolved.\n", trouble)
	}
}

func (s *BaseSupport) String() string {
	return "[" + s.name + "]"
}

// NoSupport never resolves anything.
type NoSupport struct {
	BaseSupport
}

func NewNoSupport(name string) *NoSupport {
	s := &NoSupport{}
	s.name = name
	s.resolve = func(Trouble) bool { return false }
	return s
}

func (s *NoSupport) Resolve(trouble Trouble) bool {
	return false
}

// LimitSupport resolves troubles with number < limit.
type LimitSupport struct {
	BaseSupport
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	s := &LimitSupport{limit: limit}
	s.name = name
	s.resolve = s.Resolve
	return s
}

func (s *LimitSupport) Resolve(trouble Trouble) bool {
	return trouble.Number < s.limit
}

// OddSupport resolves troubles with odd numbers.
type OddSupport struct {
	BaseSupport
}

func NewOddSupport(name string) *OddSupport {
	s := &OddSupport{}
	s.name = name
	s.resolve = s.Resolve
	return s
}

func (s *OddSupport) Resolve(trouble Trouble) bool {
	return trouble.Number%2 == 1
}

// SpecialSupport resolves only one specific trouble number.
type SpecialSupport struct {
	BaseSupport
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	s := &SpecialSupport{number: number}
	s.name = name
	s.resolve = s.Resolve
	return s
}

func (s *SpecialSupport) Resolve(trouble Trouble) bool {
	return trouble.Number == s.number
}

func main() {
	alice := NewNoSupport("Alice")
	bob := NewLimitSupport("Bob", 100)
	charlie := NewSpecialSupport("Charlie", 429)
	diana := NewLimitSupport("Diana", 200)
	elmo := NewOddSupport("Elmo")
	fred := NewLimitSupport("Fred", 300)

	// Build the chain: Alice -> Bob -> Charlie -> Diana -> Elmo -> Fred
	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	// Generate various troubles
	for i := 0; i < 500; i += 33 {
		alice.Handle(Trouble{Number: i})
	}
}
