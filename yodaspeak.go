package main

import (
	"fmt"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

type Yoda interface {
	Translate(text string) (string, error)
}

func NewYoda(key string) Yoda {
	return &YodaSpeak{
		Key: key,
		URL: "https://yoda.p.mashape.com/yoda?sentence=%s",
	}
}

type YodaSpeak struct {
	Key string
	URL string
}

func (y *YodaSpeak) Translate(text string) (string, error) {
	url := fmt.Sprintf(y.URL, url.QueryEscape(text))
	_, body, errs := gorequest.New().
		Get(url).
		Set("X-Mashape-Key", y.Key).
		End()
	if errs != nil {
		return "", errs[0]
	}

	return body, nil
}
