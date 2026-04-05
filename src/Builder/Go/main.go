// Builder: Goでも有用なパターン。Functional Optionsの代わりに、インターフェースによる
// ステップバイステップのビルダーを使い、DirectorがBuilderを駆動して文書を構築する。
package main

import (
	"fmt"
	"os"
	"strings"
)

// Builder defines the steps to construct a document.
type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

// Construct drives a Builder to build a "Greeting" document.
// This is the Director's role -- it knows the construction recipe.
func Construct(b Builder) {
	b.MakeTitle("Greeting")
	b.MakeString("一般的なあいさつ")
	b.MakeItems([]string{
		"How are you?",
		"Hello.",
		"Hi.",
	})
	b.MakeString("時間帯に応じたあいさつ")
	b.MakeItems([]string{
		"Good morning.",
		"Good afternoon.",
		"Good evening.",
	})
	b.Close()
}

// TextBuilder builds a plain-text document.
type TextBuilder struct {
	sb strings.Builder
}

func (b *TextBuilder) MakeTitle(title string) {
	b.sb.WriteString("==============================\n")
	b.sb.WriteString("『")
	b.sb.WriteString(title)
	b.sb.WriteString("』\n\n")
}

func (b *TextBuilder) MakeString(str string) {
	b.sb.WriteString("■")
	b.sb.WriteString(str)
	b.sb.WriteString("\n\n")
}

func (b *TextBuilder) MakeItems(items []string) {
	for _, s := range items {
		b.sb.WriteString("\u3000・")
		b.sb.WriteString(s)
		b.sb.WriteString("\n")
	}
	b.sb.WriteString("\n")
}

func (b *TextBuilder) Close() {
	b.sb.WriteString("==============================\n")
}

func (b *TextBuilder) Result() string {
	return b.sb.String()
}

// HTMLBuilder builds an HTML document and writes it to a file.
type HTMLBuilder struct {
	filename string
	sb       strings.Builder
}

func (b *HTMLBuilder) MakeTitle(title string) {
	b.filename = title + ".html"
	b.sb.WriteString("<!DOCTYPE html>\n")
	b.sb.WriteString("<html>\n")
	b.sb.WriteString("<head><title>")
	b.sb.WriteString(title)
	b.sb.WriteString("</title></head>\n")
	b.sb.WriteString("<body>\n")
	b.sb.WriteString("<h1>")
	b.sb.WriteString(title)
	b.sb.WriteString("</h1>\n\n")
}

func (b *HTMLBuilder) MakeString(str string) {
	b.sb.WriteString("<p>")
	b.sb.WriteString(str)
	b.sb.WriteString("</p>\n\n")
}

func (b *HTMLBuilder) MakeItems(items []string) {
	b.sb.WriteString("<ul>\n")
	for _, s := range items {
		b.sb.WriteString("<li>")
		b.sb.WriteString(s)
		b.sb.WriteString("</li>\n")
	}
	b.sb.WriteString("</ul>\n\n")
}

func (b *HTMLBuilder) Close() {
	b.sb.WriteString("</body>")
	b.sb.WriteString("</html>\n")
	if err := os.WriteFile(b.filename, []byte(b.sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	}
}

func (b *HTMLBuilder) Result() string {
	return b.filename
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go text    テキストで文書作成")
		fmt.Println("Usage: go run main.go html    HTMLファイルで文書作成")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "text":
		b := &TextBuilder{}
		Construct(b)
		fmt.Print(b.Result())
	case "html":
		b := &HTMLBuilder{}
		Construct(b)
		fmt.Printf("HTMLファイル%sが作成されました。\n", b.Result())
	default:
		fmt.Println("Usage: go run main.go text    テキストで文書作成")
		fmt.Println("Usage: go run main.go html    HTMLファイルで文書作成")
		os.Exit(0)
	}
}
