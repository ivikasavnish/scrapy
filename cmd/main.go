package main

import "github.com/ivikasavnish/scrapygo/pipeline"

func main() {
	p := pipeline.NewURLService()
	p.Link = "https://go-colly.org/"
	p.Run()
}
