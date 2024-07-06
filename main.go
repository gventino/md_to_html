package main

import(
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"net/http"
	"fmt"
)

// The actual main function from the markdow lib documentation
func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	htmlString := ""

	// Routing	
	
	// root
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)	
	
	// update
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			content := r.FormValue("md")

			// Parsing
			md := []byte(content)
			htmlBytes := mdToHTML(md)
			htmlString = string(htmlBytes[:])
		}
		
		if r.Method == http.MethodGet {
			fmt.Fprint(w, htmlString)
		}
	})

	// Putting server live
	fmt.Println("Running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
