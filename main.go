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

func (m *Movie) Info() {
	m.GetImdbInfo()
	fmt.Printf("--title \"%s\"\n", m.Title)
}

type Movie struct {
	Id          string
	Title       string
	Director    string
	Year        string
	Description string
	Genre       string
}

func main() {
	var id string

	flag.StringVar(&id, "id", "", "IMDB ID of movie. (e.g. tt1564349)")
	flag.Parse()

	if id != "" {
		m := Movie{Id: id}
		m.Info()
	}
}
