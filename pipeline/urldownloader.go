package pipeline

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/ivikasavnish/scrapygo/utils"
)

type URLService struct {
	Link    string
	Content string
	Error   error
}

func NewURLService() *URLService {
	return &URLService{}
}

func (u *URLService) Run() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {

	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		log.Println("Link found:", link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(u.Link)

	u.Content, u.Error = utils.FetchPage(u.Link)

}

func (u *URLService) Parse() {

}
