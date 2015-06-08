package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path"
)

type File struct {
	ImdbID   string
	Format   string
	FileName string
	FullPath string
}

func (f *File) IsValid() bool {
	return f.ImdbID != "" && f.Present()
}

// Present checks to see if the file is present on disk.
func (f *File) Present() bool {
	return fileExists(f.FullPath)
}

func NewFile(c *cli.Context) File {
	var fileName, format, id, fullPath string

	if len(c.Args()) == 1 {
		fullPath = c.Args()[0]
		base := path.Base(fullPath)
		format = path.Ext(fullPath)

		fileName = base[:len(base)-len(format)]
		id = c.String("id")
	}
	return File{ImdbID: id, FileName: fileName, Format: format, FullPath: fullPath}
}

func fileExists(path string) bool {
	exists := false
	if _, err := os.Stat(path); err == nil {
		exists = true
	}
	return exists
}
