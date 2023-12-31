package spiders

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type SpiderWorker struct {
	SpiderItem map[string]*Spider
}

func NewSpiderWorker() *SpiderWorker {
	return &SpiderWorker{
		SpiderItem: make(map[string]*Spider),
	}
}

func (sw *SpiderWorker) Save() {

	// try to save the spider item to a file
	out, err := json.Marshal(sw.SpiderItem)

	if err != nil {
		panic(err)
	}
	log.Println("Saving spider item to file", string(out))
	err = os.WriteFile("spider.json", out, 0777)
	if err != nil {
		log.Println(err)
	}
}

func (sw *SpiderWorker) Load() {

	// try to load the spider item from a file
	in, err := os.ReadFile("spider.json")

	if err != nil {
		panic(err)
	}
	json.Unmarshal(in, &sw.SpiderItem)
}

func (sw *SpiderWorker) Start(wg *sync.WaitGroup) {
	for _, s := range sw.SpiderItem {
		wg.Add(1)
		go s.Start(wg)
	}

}

func (sw *SpiderWorker) Stop() {

}

func (sw *SpiderWorker) Pause() {

}

func (sw *SpiderWorker) Resume() {}

func (sw *SpiderWorker) AddSpider(s *Spider) {
	// create a hash of the starting url and add it to the map
	hasher := sha256.New()
	hasher.Write([]byte(s.StartingURL))
	urlhash := hex.EncodeToString(hasher.Sum(nil))

	sw.SpiderItem[urlhash] = s
	defer sw.Save()
}
