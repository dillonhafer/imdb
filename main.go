// Package main provides an automatic way to set mp4v2 tags from IMDb
package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
)

const VERSION = "0.5.0"

var API_KEY = os.Getenv("API_KEY")

func main() {
	app := cli.NewApp()
	app.Name = "imdb-tags"
	app.Usage = "Import ID3 tags from IMDB"
	app.Version = VERSION
	app.Author = "Dillon Hafer"
	app.Email = "dh@dillonhafer.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "id, i",
			Usage: "imdb id of movie (e.g. tt1564349)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "atomic",
			Aliases: []string{"a"},
			Usage:   "Open the download page for AtomicParsley",
			Action: func(c *cli.Context) {
				webbrowser.Open("http://sourceforge.net/projects/atomicparsley/files/atomicparsley/AtomicParsley%20v0.9.0/")
			},
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search for an IMDB id by movie title",
			Action: func(c *cli.Context) {
				title := c.Args().First()
				if title != "" {
					m := SearchMovie(title)
					println("Found Possible Match:", m.Info())
				} else {
					println("No title given")
				}
			},
		},
	}

	app.Action = func(c *cli.Context) {
		CheckAtomicParsley()
		VerifyApiKey()
		file := NewFile(c)
		if file.IsValid() {
			m := FindMovie(file.ImdbID)
			t := &Tagger{Movie: m, File: file}
			t.SetTags()
		} else {
			cli.ShowAppHelp(c)
		}
	}
	app.Run(os.Args)
}

func AtomicParsleyExists() bool {
	separator := string(os.PathListSeparator)
	paths := strings.Split(os.Getenv("PATH"), separator)
	existence := false

	for _, dir := range paths {
		fullPath := filepath.Join(dir, "AtomicParsley")
		file := File{FullPath: fullPath}
		if file.Present() {
			existence = true
		}
	}
	return existence
}

func CheckAtomicParsley() {
	if !AtomicParsleyExists() {
		println("AtomicParsley is missing")
		println("You can open the download page with: `imdb-tags atomic`")
		os.Exit(1)
	}
}

func VerifyApiKey() {
	if API_KEY == "" {
		println("You must provide an API key (e.g. API_KEY=xxxxxxxx imdb-tags -i tt1564349 path/to/movie.mp4)")
		println("You can request a free one at `http://www.omdbapi.com/apikey.aspx`")
		os.Exit(1)
	}
}
