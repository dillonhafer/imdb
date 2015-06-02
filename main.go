package main

import (
	"github.com/codegangsta/cli"
	"github.com/toqueteos/webbrowser"
	"os"
	"path"
)

const VERSION = "0.0.2"

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
					println("Looking for:", title)
					m := SearchMovie(title)
					println("Possible Match:")
					println(m.Info())
				} else {
					println("No title given")
				}
			},
		},
	}

	app.Action = func(c *cli.Context) {
		var file string
		var format string
		var id string
		if len(c.Args()) == 1 {
			given_file := c.Args()[0]
			base := path.Base(given_file)
			format = path.Ext(given_file)

			file = base[:len(base)-len(format)]
			id = c.String("id")
		}

		if id != "" && file != "" && format != "" {
			m := FindMovie(id)
			t := &Tagger{Movie: m, Format: format, FilePath: file}
			t.SetTags()
		} else {
			cli.ShowAppHelp(c)
		}
	}
	app.Run(os.Args)
}
