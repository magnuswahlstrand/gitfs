package main

import (
	"fmt"
	"github.com/magnuswahlstrand/gitfs"
	"io/fs"
	"log"
)

func main() {
	gitFS, err := gitfs.New(".")
	if err != nil {
		log.Fatalf("failed to setup repository: %s", err)
	}

	err = fs.WalkDir(gitFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatalf("failed to walk dir: %s\n", err)
	}
}
