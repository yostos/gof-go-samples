// Visitor pattern is unnecessary in Go. A type switch on an interface replaces
// double dispatch (accept/visit), providing the same extensibility more directly.
package main

import "fmt"

// Entry represents a file system entry (file or directory).
type Entry interface {
	Name() string
	Size() int
}

// File is a leaf entry with a name and size.
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int    { return f.size }

func (f *File) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}

// Directory is a composite entry that contains other entries.
type Directory struct {
	name    string
	entries []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string { return d.name }

func (d *Directory) Size() int {
	total := 0
	for _, e := range d.entries {
		total += e.Size()
	}
	return total
}

func (d *Directory) Add(entry Entry) *Directory {
	d.entries = append(d.entries, entry)
	return d
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s (%d)", d.name, d.Size())
}

// ListEntries prints all entries with full paths using a type switch
// instead of the Visitor pattern's double dispatch.
func ListEntries(entry Entry, currentDir string) {
	switch e := entry.(type) {
	case *File:
		fmt.Printf("%s/%s\n", currentDir, e)
	case *Directory:
		fmt.Printf("%s/%s\n", currentDir, e)
		for _, child := range e.entries {
			ListEntries(child, currentDir+"/"+e.Name())
		}
	}
}

func main() {
	fmt.Println("Making root entries...")
	rootdir := NewDirectory("root")
	bindir := NewDirectory("bin")
	tmpdir := NewDirectory("tmp")
	usrdir := NewDirectory("usr")
	rootdir.Add(bindir).Add(tmpdir).Add(usrdir)
	bindir.Add(NewFile("vi", 10000))
	bindir.Add(NewFile("latex", 20000))
	ListEntries(rootdir, "")
	fmt.Println()

	fmt.Println("Making user entries...")
	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")
	usrdir.Add(yuki).Add(hanako).Add(tomura)
	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.java", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))
	ListEntries(rootdir, "")
}
