package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"github.com/magnuswahlstrand/gitfs"
	"io"
	"log"
)

func main() {
	gitFS, err := gitfs.New("https://github.com/magnuswahlstrand/til.git", "/src/pages/blog")
	if err != nil {
		log.Fatalf("failed to setup repository: %s", err)
	}

	for _, filename := range gitFS.Filenames {
		r, err := gitFS.Open(filename)
		if err != nil {
			log.Fatalf("failed to loop over files: %w", err)
		}
		s, err := io.ReadAll(r)
		if err != nil {
			log.Fatalf("failed to read content: %w", err)
		}
		_ = r.Close()

		out, err := glamour.RenderBytes(s, "dark")
		if err != nil {
			log.Fatalf("failed to render markdown: %w", err)
		}
		// Output markdown
		fmt.Println(string(out))
	}
}
