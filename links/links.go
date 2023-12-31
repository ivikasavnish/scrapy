package links

import "time"

//  All links will be stored in a map[string]bool

type Links struct {
	URL         string
	Parsed      bool
	LastParsed  time.Time
	LastUpdated time.Time
	Depth       int
	DomainID    string
	How         string
	Where       string
	What        string
	Why         string
	When        time.Time
	Who         string
}

type LinkStore struct {
	Links map[string]*Links
}

func NewLinkStore() *LinkStore {
	return &LinkStore{
		Links: make(map[string]*Links),
	}
}

func (ls *LinkStore) AddLink(link *Links) {
	ls.Links[link.URL] = link
}

func (ls *LinkStore) Run(url string) {

}
