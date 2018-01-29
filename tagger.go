package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type Tagger struct {
	Movie Movie
	File  File
}

func (t *Tagger) TmpFileName() string {
	return fmt.Sprintf("%s%s%s", t.File.FileName, t.TempID(), t.File.Format)
}

func (t *Tagger) FullTmpFileName() string {
	fullPath := filepath.Dir(t.File.FullPath)
	return filepath.Join(fullPath, t.TmpFileName())
}

func (t *Tagger) TempID() string {
	dir := filepath.Dir(t.File.FullPath)
	var fileID string
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		r, _ := regexp.Compile("-temp-([0-9]+)")
		if r.MatchString(f.Name()) {
			fileID = r.FindString(f.Name())
		}
	}
	return fileID
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

		resp, err := check.Get(t.Movie.ArtworkURL) // add a filter to check redirect
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
	fileArgs := []string{t.File.FullPath}
	args := append(fileArgs, t.Movie.ParsleyFlags()...)

	if t.Movie.HasArtwork() {
		artwork := []string{"--artwork", "REMOVE_ALL", "--artwork", "artwork.jpg"}
		args = append(args, artwork...)
	} else {
		println("Could not find artwork")
	}

	if t.Movie.IsValid() {
		pwd, _ := os.Getwd()
		ap := filepath.Join(pwd, AtomicParsley)
		out, err := exec.Command(ap, args...).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	} else {
		println("Could not find IMDB info")
	}
}

func (t *Tagger) CleanupCommand() {
	println("Cleaning up tmp files")
	os.Remove("artwork.jpg")
	os.Rename(t.FullTmpFileName(), t.File.FullPath)
}

func (t *Tagger) SetTags() {
	defer t.CleanupCommand()
	t.GetArtwork()
	t.AtomicCommand()
}
