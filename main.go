package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path"
)

const VERSION = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "imdb"
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
			println(id)
			println(file)
			println(format)
			m := FindMovie(id)
			t := &Tagger{Movie: m, Format: format, FilePath: file}
			t.SetTags()
		} else {
			cli.ShowAppHelp(c)
		}
	}
	app.Run(os.Args)
}
