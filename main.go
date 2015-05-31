package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

type Movie struct {
	Id       string
	Title    string
	Director string
	Year     string
	Plot     string
	Genre    string
	Poster   string
}

type Tagger struct {
	Movie    Movie
	FilePath string
	Format   string
}

func (m *Movie) GetImdbInfo() {
	res, err := http.Get(m.ApiUrl())
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		panic(err.Error())
	}
}

func (m *Movie) ApiUrl() string {
	url := "http://www.omdbapi.com/?i=%s&plot=short&r=json"
	return fmt.Sprintf(url, m.Id)
}

func (m *Movie) ParsleyFlags() []string {
	return []string{"--title", m.Title, "--artist", m.Director, "--year", m.Year, "--description", m.Plot, "--genre", m.Genre}
}

func FindMovie(id string) Movie {
	m := Movie{Id: id}
	m.GetImdbInfo()
	return m
}

func (t *Tagger) TempId() string {
	var file_id string
	files, _ := ioutil.ReadDir("./")
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
	file_path := fmt.Sprintf("%s.%s", t.FilePath, t.Format)
	file_args := []string{file_path}
	args := append(file_args, t.Movie.ParsleyFlags()...)

	if t.Movie.Poster != "" {
		artwork := []string{"--artwork", "REMOVE_ALL", "--artwork", "artwork.jpg"}
		args = append(args, artwork...)
	}

	_, err := exec.Command("AtomicParsley", args...).Output()
	return err
}

func (t *Tagger) CleanupCommand() error {
	old_file := fmt.Sprintf("%s%s.%s", t.FilePath, t.TempId(), t.Format)
	new_file := fmt.Sprintf("%s.%s", t.FilePath, t.Format)
	os.Remove("artwork.jpg")
	return os.Rename(old_file, new_file)
}

func (t *Tagger) SetTags() {
	t.GetArtwork()
	fmt.Printf("Setting tags...\n")
	t.AtomicCommand()
	fmt.Printf("Cleaning up tmp files...\n")
	t.CleanupCommand()
}

func main() {
	var id string
	var file_path string
	var format string
	flag.StringVar(&id, "id", "", "IMDB ID of movie (e.g. tt1564349)")
	flag.StringVar(&file_path, "file", "", "Path to video file")
	flag.StringVar(&format, "format", "m4v", "File format of video file (defaults to m4v)")
	flag.Parse()

	if id != "" && file_path != "" && format != "" {
		m := FindMovie(id)
		t := &Tagger{Movie: m, Format: format, FilePath: file_path}
		t.SetTags()
	} else {
		flag.Usage()
	}
}
