package main

import "github.com/magnuswahlstrand/gitfs"

func main() {
	gitFS, err := gitfs.NewGit("https://github.com/magnuswahlstrand/til.git", "/src/pages/blog")
	//if err != nil {
	//	log.Fatalf("failed to setup repository: %s", err)
	//}
	//
	//for _, filename := range gitFS.Filenames {
	//	r, err := gitFS.Open(filename)
	//	if err != nil {
	//		log.Fatalf("failed to loop over files: %w", err)
	//	}
	//	s, err := io.ReadAll(r)
	//	if err != nil {
	//		log.Fatalf("failed to read content: %w", err)
	//	}
	//	_ = r.Close()
	//
	//	out, err := glamour.RenderBytes(s, "dark")
	//	if err != nil {
	//		log.Fatalf("failed to render markdown: %w", err)
	//	}
	//	// Output markdown
	//	fmt.Println(string(out))
	//}
}
