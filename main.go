package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func (m *Movie) ParsleyFlags() {
	fmt.Printf("--title \"%s\" --artist \"%s\" --year \"%s\" --description \"%s\" --genre \"%s\"", m.Title, m.Director, m.Year, m.Plot, m.Genre)
}

type Movie struct {
	Id       string
	Title    string
	Director string
	Year     string
	Plot     string
	Genre    string
}

func FindMovie(id string) Movie {
	m := Movie{Id: id}
	m.GetImdbInfo()
	return m
}

func main() {
	var id string
	flag.StringVar(&id, "id", "", "IMDB ID of movie. (e.g. tt1564349)")
	flag.Parse()

	if id != "" {
		m := FindMovie(id)
		m.ParsleyFlags()
	}
}
