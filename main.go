package main

import (
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"os"
	"path"
	"strings"
)

const VERSION = "0.1.0"

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
					println("Found Possible Match:\n%s", m.Info())
				} else {
					println("No title given")
				}
			},
		},
	}

	app.Action = func(c *cli.Context) {
		CheckAtomicParsley()
		params := GetParams(c)
		if params.IsValid() {
			m := FindMovie(params.id)
			t := &Tagger{Movie: m, Format: params.format, FilePath: params.file}
			t.SetTags()
		} else {
			cli.ShowAppHelp(c)
		}
	}
	app.Run(os.Args)
}

type Params struct {
	id     string
	format string
	file   string
}

func (p *Params) IsValid() bool {
	return p.id != "" && p.format != "" && p.file != ""
}

func AtomicParsleyExists() bool {
	paths := strings.Split(os.Getenv("PATH"), ":")
	existence := false

	for _, each := range paths {
		path := path.Join(each, "AtomicParsley")
		if _, err := os.Stat(path); err == nil {
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

func GetParams(c *cli.Context) Params {
	var file, format, id string

	if len(c.Args()) == 1 {
		given_file := c.Args()[0]
		base := path.Base(given_file)
		format = path.Ext(given_file)

		file = base[:len(base)-len(format)]
		id = c.String("id")
	}
	return Params{id: id, file: file, format: format}
}
