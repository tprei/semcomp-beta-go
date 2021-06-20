package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
)

type engine struct {
	client *http.Client

	requests  chan *http.Request
	responses chan *http.Response

	done chan bool
}

func cleanString(s string) string {
	return strings.ToLower(strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, s))
}

func (eng *engine) worker() {
	for job := range eng.requests {
		if resp, err := eng.client.Do(job); err != nil {
			log.Println("[ERROR]", err)
		} else {
			eng.responses <- resp
		}
	}
}

func scrape(numWorkers int) {
	f, _ := os.Open("pokemon.csv")
	reader := csv.NewReader(bufio.NewReader(f))
	records, _ := reader.ReadAll()

	eng := engine{
		client:    http.DefaultClient,
		requests:  make(chan *http.Request, len(records)),
		responses: make(chan *http.Response, len(records)),
	}

	for _, record := range records {
		pokemonName := cleanString(record[1])
		url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

		if req, err := http.NewRequest(http.MethodGet, url, nil); err != nil {
			log.Fatal(err)
		} else {
			eng.requests <- req
		}

	}

	close(eng.requests)

	for i := 0; i < numWorkers; i++ {
		go eng.worker()
	}

	for i := 0; i < len(records); i++ {
		<-eng.responses
		//resp := <-eng.responses
		//log.Println(resp.StatusCode, resp.Request.URL)
	}
}

func main() {
	num := 50
	initial := time.Now()
	scrape(num)
	elapsed := time.Now().Sub(initial).Seconds()
	fmt.Printf("elapsed was %v with %v workers\n", elapsed, num)
}
