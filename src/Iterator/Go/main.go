// Iterator pattern: in Go 1.23, iter.Seq replaces explicit Iterator classes.
// BookShelf.All() returns an iter.Seq[Book] that works with range-over-function.
package main

import (
	"fmt"
	"iter"
)

// Book represents a book with a name.
type Book struct {
	Name string
}

// BookShelf holds a collection of books.
type BookShelf struct {
	books []Book
}

// Append adds a book to the shelf.
func (bs *BookShelf) Append(book Book) {
	bs.books = append(bs.books, book)
}

// All returns an iter.Seq[Book] for use with range-over-function.
// This is the idiomatic Go 1.23 way to provide iteration.
func (bs *BookShelf) All() iter.Seq[Book] {
	return func(yield func(Book) bool) {
		for _, book := range bs.books {
			if !yield(book) {
				return
			}
		}
	}
}

func main() {
	bookShelf := &BookShelf{}
	bookShelf.Append(Book{Name: "Around the World in 80 Days"})
	bookShelf.Append(Book{Name: "Bible"})
	bookShelf.Append(Book{Name: "Cinderella"})
	bookShelf.Append(Book{Name: "Daddy-Long-Legs"})

	// Use range-over-function with iter.Seq (Go 1.23)
	for book := range bookShelf.All() {
		fmt.Println(book.Name)
	}
	fmt.Println()

	// The same iteration can be used multiple times
	for book := range bookShelf.All() {
		fmt.Println(book.Name)
	}
	fmt.Println()
}
