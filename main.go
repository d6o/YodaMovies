package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	var key string
	flag.StringVar(&key, "key", "", "Your Mashape Key")
	flag.Parse()

	if len(key) == 0 {
		fmt.Println("Please provide a key. Create your free key at https://market.mashape.com/signup")
		return
	}
	quote, err := NewMoviesQuotes(key).Get()
	if err != nil {
		panic(err)
	}

	translatedQuote, err := NewYoda(key).Translate(quote.Text)
	if err != nil {
		panic(err)
	}

	green := color.New(color.FgGreen).SprintFunc()
	darkGreen := color.New(color.FgHiGreen).SprintFunc()

	fmt.Println("Quote: " + quote.Text)
	fmt.Println("Movie: " + quote.Author)
	fmt.Printf("%s: %s \n", darkGreen("Yoda"), green(translatedQuote))
}
