package spiders

import (
	"sync"

	"github.com/ivikasavnish/scrapygo/pipeline"
)

type Spider struct {
	StartingURL    string
	MaxDepth       int
	MaxWorkers     int
	SurfaceRefresh int
	LinkMash       map[string]bool
	Using          string
	Extractors     string
}

func (s *Spider) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	flow := pipeline.NewURLService()
	flow.Link = s.StartingURL
	flow.Run()

}
