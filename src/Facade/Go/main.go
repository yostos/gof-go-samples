// Facade pattern in Go: PageMaker provides a simple API that hides the
// complexity of reading a database and generating HTML output.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// database reads a key=value properties file and returns a map.
func loadDatabase(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	props := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if idx := strings.Index(line, "="); idx >= 0 {
			props[line[:idx]] = line[idx+1:]
		}
	}
	return props, scanner.Err()
}

// htmlWriter writes simple HTML to an io.Writer.
type htmlWriter struct {
	w io.Writer
}

func (h *htmlWriter) title(title string) {
	fmt.Fprint(h.w, "<!DOCTYPE html>")
	fmt.Fprint(h.w, "<html>")
	fmt.Fprint(h.w, "<head>")
	fmt.Fprintf(h.w, "<title>%s</title>", title)
	fmt.Fprint(h.w, "</head>")
	fmt.Fprint(h.w, "<body>")
	fmt.Fprintln(h.w)
	fmt.Fprintf(h.w, "<h1>%s</h1>", title)
	fmt.Fprintln(h.w)
}

func (h *htmlWriter) paragraph(msg string) {
	fmt.Fprintf(h.w, "<p>%s</p>", msg)
	fmt.Fprintln(h.w)
}

func (h *htmlWriter) link(href, caption string) {
	h.paragraph(fmt.Sprintf("<a href=\"%s\">%s</a>", href, caption))
}

func (h *htmlWriter) mailto(addr, name string) {
	h.link("mailto:"+addr, name)
}

func (h *htmlWriter) close() {
	fmt.Fprint(h.w, "</body>")
	fmt.Fprint(h.w, "</html>")
	fmt.Fprintln(h.w)
}

// makeWelcomePage is the facade function that orchestrates everything.
func makeWelcomePage(mailaddr, filename string) error {
	maildata, err := loadDatabase("maildata.txt")
	if err != nil {
		return fmt.Errorf("loading maildata: %w", err)
	}

	username, ok := maildata[mailaddr]
	if !ok {
		return fmt.Errorf("mail address %s not found", mailaddr)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating %s: %w", filename, err)
	}
	defer f.Close()

	hw := &htmlWriter{w: f}
	hw.title(username + "'s web page")
	hw.paragraph("Welcome to " + username + "'s web page!")
	hw.paragraph("Nice to meet you!")
	hw.mailto(mailaddr, username)
	hw.close()

	fmt.Printf("%s is created for %s (%s)\n", filename, mailaddr, username)
	return nil
}

func main() {
	if err := makeWelcomePage("hyuki@example.com", "welcome.html"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
