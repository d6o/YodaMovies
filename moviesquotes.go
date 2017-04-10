package main

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type MoviesQuotes interface {
	Get() (Quote, error)
}

type Quote struct {
	Text   string
	Author string
}

func NewMoviesQuotes(key string) MoviesQuotes {
	return &RandomFamousQuotes{
		URL: "https://andruxnet-random-famous-quotes.p.mashape.com/?cat=movies",
		Key: key,
	}
}

type RandomFamousQuotes struct {
	URL string
	Key string
}

type RandomFamousQuotesResponse struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (r *RandomFamousQuotes) Get() (Quote, error) {
	q := Quote{}
	_, body, errs := gorequest.New().
		Get(r.URL).
		Set("X-Mashape-Key", r.Key).
		End()
	if errs != nil {
		return q, errs[0]
	}

	resp := RandomFamousQuotesResponse{}
	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return q, err
	}

	q.Author = resp.Author
	q.Text = resp.Quote

	return q, nil
}
