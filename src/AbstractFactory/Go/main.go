// AbstractFactory: Goではクラス継承がないため、インターフェースとファクトリ関数で十分。
// 抽象基底クラスの階層は不要で、シンプルなインターフェースと関数値で同じ問題を解決する。
package main

import (
	"fmt"
	"os"
	"strings"
)

// Link represents a hyperlink.
type Link struct {
	Caption string
	URL     string
}

// Tray holds child items under a caption. Items can be *Link or *Tray.
type Tray struct {
	Caption string
	Items   []any
}

// Page represents an HTML page with a title, author, and content items.
type Page struct {
	Title   string
	Author  string
	Content []any
}

// Add appends an item to the tray.
func (t *Tray) Add(item any) {
	t.Items = append(t.Items, item)
}

// Add appends an item to the page.
func (p *Page) Add(item any) {
	p.Content = append(p.Content, item)
}

// PageRenderer defines how a page renders its HTML.
// This is the only behavior that varies between "list" and "div" styles.
type PageRenderer struct {
	RenderLink func(link *Link) string
	RenderTray func(tray *Tray, renderItem func(any) string) string
	RenderPage func(page *Page, renderItem func(any) string) string
}

// renderItem dispatches to the correct renderer based on type.
func (pr *PageRenderer) renderItem(item any) string {
	switch v := item.(type) {
	case *Link:
		return pr.RenderLink(v)
	case *Tray:
		return pr.RenderTray(v, pr.renderItem)
	default:
		return ""
	}
}

// Output writes the page HTML to a file.
func (pr *PageRenderer) Output(page *Page, filename string) error {
	html := pr.RenderPage(page, pr.renderItem)
	if err := os.WriteFile(filename, []byte(html), 0644); err != nil {
		return err
	}
	fmt.Printf("%s を作成しました。\n", filename)
	return nil
}

// ListRenderer returns a PageRenderer that uses <ul>/<li> markup.
func ListRenderer() *PageRenderer {
	return &PageRenderer{
		RenderLink: func(link *Link) string {
			return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", link.URL, link.Caption)
		},
		RenderTray: func(tray *Tray, renderItem func(any) string) string {
			var sb strings.Builder
			sb.WriteString("<li>\n")
			sb.WriteString(tray.Caption)
			sb.WriteString("\n<ul>\n")
			for _, item := range tray.Items {
				sb.WriteString(renderItem(item))
			}
			sb.WriteString("</ul>\n")
			sb.WriteString("</li>\n")
			return sb.String()
		},
		RenderPage: func(page *Page, renderItem func(any) string) string {
			var sb strings.Builder
			sb.WriteString("<!DOCTYPE html>\n")
			sb.WriteString("<html><head><title>")
			sb.WriteString(page.Title)
			sb.WriteString("</title></head>\n")
			sb.WriteString("<body>\n")
			sb.WriteString("<h1>")
			sb.WriteString(page.Title)
			sb.WriteString("</h1>\n")
			sb.WriteString("<ul>\n")
			for _, item := range page.Content {
				sb.WriteString(renderItem(item))
			}
			sb.WriteString("</ul>\n")
			sb.WriteString("<hr><address>")
			sb.WriteString(page.Author)
			sb.WriteString("</address>\n")
			sb.WriteString("</body></html>\n")
			return sb.String()
		},
	}
}

// DivRenderer returns a PageRenderer that uses <div> markup.
func DivRenderer() *PageRenderer {
	return &PageRenderer{
		RenderLink: func(link *Link) string {
			return fmt.Sprintf("<div class=\"LINK\"><a href=\"%s\">%s</a></div>\n", link.URL, link.Caption)
		},
		RenderTray: func(tray *Tray, renderItem func(any) string) string {
			var sb strings.Builder
			sb.WriteString("<p><b>")
			sb.WriteString(tray.Caption)
			sb.WriteString("</b></p>\n")
			sb.WriteString("<div class=\"TRAY\">")
			for _, item := range tray.Items {
				sb.WriteString(renderItem(item))
			}
			sb.WriteString("</div>\n")
			return sb.String()
		},
		RenderPage: func(page *Page, renderItem func(any) string) string {
			var sb strings.Builder
			sb.WriteString("<!DOCTYPE html>\n")
			sb.WriteString("<html><head><title>")
			sb.WriteString(page.Title)
			sb.WriteString("</title><style>\n")
			sb.WriteString("div.TRAY { padding:0.5em; margin-left:5em; border:1px solid black; }\n")
			sb.WriteString("div.LINK { padding:0.5em; background-color: lightgray; }\n")
			sb.WriteString("</style></head><body>\n")
			sb.WriteString("<h1>")
			sb.WriteString(page.Title)
			sb.WriteString("</h1>\n")
			for _, item := range page.Content {
				sb.WriteString(renderItem(item))
			}
			sb.WriteString("<hr><address>")
			sb.WriteString(page.Author)
			sb.WriteString("</address>\n")
			sb.WriteString("</body></html>\n")
			return sb.String()
		},
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <filename.html> <list|div>")
		fmt.Println("Example 1: go run main.go list.html list")
		fmt.Println("Example 2: go run main.go div.html div")
		os.Exit(0)
	}

	filename := os.Args[1]
	style := os.Args[2]

	// Select renderer by name -- no reflection or class loading needed.
	var renderer *PageRenderer
	switch style {
	case "list":
		renderer = ListRenderer()
	case "div":
		renderer = DivRenderer()
	default:
		fmt.Fprintf(os.Stderr, "Unknown style: %s (use 'list' or 'div')\n", style)
		os.Exit(1)
	}

	// Blog
	blog1 := &Link{Caption: "Blog 1", URL: "https://example.com/blog1"}
	blog2 := &Link{Caption: "Blog 2", URL: "https://example.com/blog2"}
	blog3 := &Link{Caption: "Blog 3", URL: "https://example.com/blog3"}

	blogTray := &Tray{Caption: "Blog Site"}
	blogTray.Add(blog1)
	blogTray.Add(blog2)
	blogTray.Add(blog3)

	// News
	news1 := &Link{Caption: "News 1", URL: "https://example.com/news1"}
	news2 := &Link{Caption: "News 2", URL: "https://example.com/news2"}
	news3 := &Tray{Caption: "News 3"}
	news3.Add(&Link{Caption: "News 3 (US)", URL: "https://example.com/news3us"})
	news3.Add(&Link{Caption: "News 3 (Japan)", URL: "https://example.com/news3jp"})

	newsTray := &Tray{Caption: "News Site"}
	newsTray.Add(news1)
	newsTray.Add(news2)
	newsTray.Add(news3)

	// Page
	page := &Page{Title: "Blog and News", Author: "Hiroshi Yuki"}
	page.Add(blogTray)
	page.Add(newsTray)

	if err := renderer.Output(page, filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
