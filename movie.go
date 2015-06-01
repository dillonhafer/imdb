package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	Id          string
	Title       string `json:"Title"`
	Director    string `json:"Director"`
	Year        string `json:"Year"`
	Description string `json:"Plot"`
	Genre       string `json:"Genre"`
	ArtworkUrl  string `json:"Poster"`
	ImdbId      string `json:"imdbID"`
}

func (m *Movie) GetImdbInfo(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		panic(err.Error())
	}
}

func (m *Movie) ApiUrl() string {
	url := "http://www.omdbapi.com/?i=%s&plot=short&r=json"
	return fmt.Sprintf(url, m.Id)
}

func (m *Movie) SearchApiUrl() string {
	url := "http://www.omdbapi.com/?t=%s&plot=short&r=json"
	return fmt.Sprintf(url, m.Title)
}

func (m *Movie) ParsleyFlags() []string {
	return []string{"--title", m.Title, "--artist", m.Director, "--year", m.Year, "--description", m.Description, "--genre", m.Genre}
}

func (m *Movie) Info() string {
	return fmt.Sprintf("  Title: %s\n  Director: %s\n  Year: %s\n  Plot: %s\n  IMDB ID: %s", m.Title, m.Director, m.Year, m.Description, m.ImdbId)
}

func (m *Movie) HasArtwork() bool {
	return m.ArtworkUrl != ""
}

func FindMovie(id string) Movie {
	m := Movie{Id: id}
	m.GetImdbInfo(m.ApiUrl())
	return m
}

func SearchMovie(title string) Movie {
	m := Movie{Title: title}
	m.GetImdbInfo(m.SearchApiUrl())
	return m
}
