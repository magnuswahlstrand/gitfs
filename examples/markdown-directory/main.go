package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/glamour"
	"github.com/magnuswahlstrand/gitfs"
	"io"
	"io/fs"
	"log"
)

func main() {
	repo := flag.String("repo", "", "URL to a git repo, can be a local file")
	directory := flag.String("path", "", "path inside the repository")
	flag.Parse()
	if *repo == "" {
		log.Fatalln("-repo is required")
	}

	if *directory == "" {
		log.Fatalln("-path is required")
	}

	gitFS, err := gitfs.New(*repo)
	if err != nil {
		log.Fatalf("failed to setup repository: %s", err)
	}

	err = fs.WalkDir(gitFS, *directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		f, err := gitFS.Open(path)
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		content, err := io.ReadAll(f)
		if err != nil {
			log.Fatalf("failed to read content: %s", err)
		}

		out, err := glamour.RenderBytes(content, "dark")
		if err != nil {
			log.Fatalf("failed to render markdown: %s", err)
		}

		// output markdown
		fmt.Println(string(out))

		return nil
	})
}
