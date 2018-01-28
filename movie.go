package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Movie struct {
	ImdbID      string `json:"imdbID"`
	Title       string `json:"Title"`
	Director    string `json:"Director"`
	Year        string `json:"Year"`
	Description string `json:"Plot"`
	Genre       string `json:"Genre"`
	ArtworkURL  string `json:"Poster"`
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

func (m *Movie) APIURL() string {
	query := fmt.Sprintf("i=%s&plot=short&r=json", m.ImdbID)
	return BaseAPI(query)
}

func (m *Movie) SearchAPIURL() string {
	query := fmt.Sprintf("t=%s&plot=short&r=json", url.QueryEscape(m.Title))
	return BaseAPI(query)
}

func BaseAPI(q string) string {
	return fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&%s", API_KEY, q)
}

func (m *Movie) ParsleyFlags() []string {
	return []string{"--title", m.Title, "--artist", m.Director, "--year", m.Year, "--description", m.Description, "--genre", m.Genre}
}

func (m *Movie) Info() string {
	return fmt.Sprintf("\n  Title: %s\n  Director: %s\n  Year: %s\n  Plot: %s\n  IMDB ID: %s", m.Title, m.Director, m.Year, m.Description, m.ImdbID)
}

func (m *Movie) HasArtwork() bool {
	return m.ArtworkURL != "" && m.ArtworkURL != "N/A"
}

func (m *Movie) IsValid() bool {
	return m.Title != "" && m.Director != "" && m.Year != "" && m.Description != "" && m.Genre != ""
}

func FindMovie(id string) Movie {
	m := Movie{ImdbID: id}
	m.GetImdbInfo(m.APIURL())
	return m
}

func SearchMovie(title string) Movie {
	m := Movie{Title: title}
	m.GetImdbInfo(m.SearchAPIURL())
	return m
}
