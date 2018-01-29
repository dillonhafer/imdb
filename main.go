// Package main provides an automatic way to set mp4v2 tags from IMDb
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

const VERSION = "0.6.0"

var API_KEY = os.Getenv("API_KEY")
var AtomicParsley = "atomic-parsley"

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
		VerifyApiKey()

		ExtractAtomicParsley()
		defer RemoveAtomicParsley()

		file := NewFile(c)
		if file.IsValid() {
			m, err := FindMovie(file.ImdbID)
			if err != nil {
				fmt.Print(err)
				return
			}

			t := &Tagger{Movie: m, File: file}
			t.SetTags()
		} else {
			cli.ShowAppHelp(c)
		}
	}
	app.Run(os.Args)
}

func ExtractAtomicParsley() {
	var assetFolder string
	switch os := runtime.GOOS; os {
	case "windows":
		assetFolder = "bin/windows/AtomicParsley.exe"
		AtomicParsley = AtomicParsley + ".exe"
	case "darwin":
		assetFolder = "bin/osx/AtomicParsley"
	case "linux":
		assetFolder = "bin/linux/AtomicParsley"
	default:
		fmt.Printf("Your operating system is not supported:", os)
	}

	if assetFolder == "" {
		os.Exit(1)
	}

	data, err := Asset(assetFolder)
	if err != nil {
		println("Could not extract AtomicParsley.")
		os.Exit(1)
	}

	err = ioutil.WriteFile(AtomicParsley, data, 0744)
	if err != nil {
		println("Could create AtomicParsley file")
		os.Exit(1)
	}
}

func RemoveAtomicParsley() {
	err := os.Remove(AtomicParsley)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func VerifyApiKey() {
	if API_KEY == "" {
		switch os := runtime.GOOS; os {
		case "windows":
			println("You must provide an API key (e.g. \nset API_KEY=xxxxxxxx\nimdb-tags -i tt1564349 path/to/movie.mp4)")
		default:
			println("You must provide an API key (e.g. API_KEY=xxxxxxxx imdb-tags -i tt1564349 path/to/movie.mp4)")
		}

		println("You can request a free one at `http://www.omdbapi.com/apikey.aspx`")
		os.Exit(1)
	}
}
