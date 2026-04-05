// Composite pattern in Go: uses an Entry interface to treat File and Directory
// uniformly, with recursive listing via a shared interface method.
package main

import "fmt"

// Entry is the component interface shared by files and directories.
type Entry interface {
	Name() string
	Size() int
	printList(prefix string)
}

// PrintList prints the full listing starting from the given entry.
func PrintList(e Entry) {
	e.printList("")
}

// entryString returns the standard "name (size)" representation.
func entryString(e Entry) string {
	return fmt.Sprintf("%s (%d)", e.Name(), e.Size())
}

// File is a leaf node.
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int    { return f.size }

func (f *File) printList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, entryString(f))
}

// Directory is a composite node that contains entries.
type Directory struct {
	name     string
	children []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string { return d.name }

func (d *Directory) Size() int {
	total := 0
	for _, child := range d.children {
		total += child.Size()
	}
	return total
}

func (d *Directory) Add(entry Entry) *Directory {
	d.children = append(d.children, entry)
	return d
}

func (d *Directory) printList(prefix string) {
	fmt.Printf("%s/%s\n", prefix, entryString(d))
	for _, child := range d.children {
		child.printList(prefix + "/" + d.name)
	}
}

func main() {
	fmt.Println("Making root entries...")
	rootdir := NewDirectory("root")
	bindir := NewDirectory("bin")
	tmpdir := NewDirectory("tmp")
	usrdir := NewDirectory("usr")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)
	bindir.Add(NewFile("vi", 10000))
	bindir.Add(NewFile("latex", 20000))
	PrintList(rootdir)
	fmt.Println()

	fmt.Println("Making user entries...")
	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")
	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)
	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.java", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))
	PrintList(rootdir)
}
