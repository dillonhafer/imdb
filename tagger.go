package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

type Tagger struct {
	Movie Movie
	File  File
}

func (t *Tagger) TmpFileName() string {
	return fmt.Sprintf("%s%s%s", t.File.FileName, t.TempId(), t.File.Format)
}

func (t *Tagger) TempId() string {
	var file_id string
	pwd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(pwd)
	for _, f := range files {
		r, _ := regexp.Compile("-temp-([0-9]+)")
		if r.MatchString(f.Name()) {
			file_id = r.FindString(f.Name())
		}
	}
	return file_id
}

func (t *Tagger) GetArtwork() {
	if t.Movie.HasArtwork() {
		println("Downloading artwork")
		file, err := os.Create("artwork.jpg")
		defer file.Close()

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}

		resp, err := check.Get(t.Movie.ArtworkUrl) // add a filter to check redirect
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer resp.Body.Close()
		io.Copy(file, resp.Body)

		if err != nil {
			panic(err)
		}
	}
}

func (t *Tagger) AtomicCommand() {
	file_args := []string{t.File.FullPath}
	args := append(file_args, t.Movie.ParsleyFlags()...)

	if t.Movie.HasArtwork() {
		artwork := []string{"--artwork", "REMOVE_ALL", "--artwork", "artwork.jpg"}
		args = append(args, artwork...)
	} else {
		println("Could not find artwork")
	}

	if t.Movie.IsValid() {
		println("Setting tags")
		exec.Command("AtomicParsley", args...).Output()
	} else {
		println("Could not find IMDB info")
	}
}

func (t *Tagger) CleanupCommand() {
	println("Cleaning up tmp files")
	os.Remove("artwork.jpg")
	os.Rename(t.TmpFileName(), t.File.FullPath)
}

func (t *Tagger) SetTags() {
	t.GetArtwork()
	t.AtomicCommand()
	t.CleanupCommand()
}
