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
	Movie    Movie
	FilePath string
	Format   string
}

func (t *Tagger) FullFileName() string {
	return fmt.Sprintf("%s%s", t.FilePath, t.Format)
}

func (t *Tagger) TmpFileName() string {
	return fmt.Sprintf("%s%s%s", t.FilePath, t.TempId(), t.Format)
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
	if t.Movie.Poster != "" {
		fmt.Printf("Downloading artwork...\n")
		file, err := os.Create("artwork.jpg")
		defer file.Close()

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}

		resp, err := check.Get(t.Movie.Poster) // add a filter to check redirect
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

func (t *Tagger) AtomicCommand() error {
	file_args := []string{t.FullFileName()}
	args := append(file_args, t.Movie.ParsleyFlags()...)

	if t.Movie.Poster != "" {
		artwork := []string{"--artwork", "REMOVE_ALL", "--artwork", "artwork.jpg"}
		args = append(args, artwork...)
	}

	_, err := exec.Command("AtomicParsley", args...).Output()
	return err
}

func (t *Tagger) CleanupCommand() {
	os.Remove("artwork.jpg")
	os.Rename(t.TmpFileName(), t.FullFileName())
}

func (t *Tagger) SetTags() {
	t.GetArtwork()
	fmt.Printf("Setting tags...\n")
	t.AtomicCommand()
	fmt.Printf("Cleaning up tmp files...\n")
	t.CleanupCommand()
}
