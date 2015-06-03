package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path"
)

type File struct {
	ImdbId   string
	Format   string
	FileName string
	FullPath string
}

func (f *File) IsValid() bool {
	return f.ImdbId != "" && f.Present()
}

// Is the file is present on disk
func (f *File) Present() bool {
	return file_exists(f.FullPath)
}

func NewFile(c *cli.Context) File {
	var file_name, format, id, full_path string

	if len(c.Args()) == 1 {
		full_path = c.Args()[0]
		base := path.Base(full_path)
		format = path.Ext(full_path)

		file_name = base[:len(base)-len(format)]
		id = c.String("id")
	}
	return File{ImdbId: id, FileName: file_name, Format: format, FullPath: full_path}
}

func file_exists(path string) bool {
	exists := false
	if _, err := os.Stat(path); err == nil {
		exists = true
	}
	return exists
}
