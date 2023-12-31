package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ivikasavnish/scrapygo/spiders"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

var (
	counterchan = make(chan int, 1000)
)

func Run(worker *spiders.SpiderWorker) {
	for {
		time.Sleep(1 * time.Minute)
		worker.Save()

	}
}

func main() {
	worker := spiders.NewSpiderWorker()

	worker.Load()
	var wg sync.WaitGroup

	worker.Start(&wg)
	wg.Wait()
	// Loadfromcsv()

}

func Loadfromcsv() {
	worker := spiders.NewSpiderWorker()
	// Load the CSV file.
	f, err := os.Open("sites.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new reader.
	reader := csv.NewReader(f)

	// Read all lines.
	var lines [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	// Display all lines.
	for _, line := range lines {
		fmt.Println(line)
		worker.AddSpider(&spiders.Spider{
			StartingURL: line[1],
			MaxDepth:    2,
			MaxWorkers:  2,
		})

	}
}
