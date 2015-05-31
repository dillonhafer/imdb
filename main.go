package main

import (
	"flag"
)

func main() {
	var id string
	var file_path string
	var format string
	flag.StringVar(&id, "id", "", "IMDB ID of movie (e.g. tt1564349)")
	flag.StringVar(&file_path, "file", "", "Path to video file")
	flag.StringVar(&format, "format", "m4v", "File format of video file (defaults to m4v)")
	flag.Parse()

	if id != "" && file_path != "" && format != "" {
		m := FindMovie(id)
		t := &Tagger{Movie: m, Format: format, FilePath: file_path}
		t.SetTags()
	} else {
		flag.Usage()
	}
}
