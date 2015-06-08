package main

import (
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"os"
	"path"
	"strings"
)

const VERSION = "0.4.2"

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
	paths := strings.Split(os.Getenv("PATH"), ":")
	existence := false

	for _, each := range paths {
		path := path.Join(each, "AtomicParsley")
		file := File{FullPath: path}
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
