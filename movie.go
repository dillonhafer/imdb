package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
