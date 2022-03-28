package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Film struct {
	Characters []string
	Title      string
	Url        string
}

func main() {
	resp, err := http.Get("https://swapi.dev/api/films/1")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var film Film
	err = json.Unmarshal(body, &film)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("%v, %v\n", film.Title, film.Url)

	// The functions below do the same thing in different ways
	//WaitGroupAndMutex(film)
	Channel(film)
}

func WaitGroupAndMutex(film Film) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, url := range film.Characters {
		wg.Add(1)
		go FetchAndPrint(url, &wg, &mutex) // must pass pointer - passing a copy is useless!
	}

	wg.Wait()
}

var count int = 0

func FetchAndPrint(url string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("\t%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	strbody := string(body)

	mutex.Lock()
	defer mutex.Unlock()
	count += 1
	fmt.Printf("%v\t[%v]\t%v:\t%v...\n", count, time.Now().Format("2006-01-02T15:04:05.999"), url, strbody[:40])
}

func Channel(film Film) {
	ch := make(chan Result)
	for _, url := range film.Characters {
		go Fetch(url, ch)
	}

	count := 0
	for range film.Characters {
		out := <-ch
		printstr := out.Value
		if printstr == "" {
			printstr = out.Error
		}
		printlen := len(printstr)
		if printlen > 40 {
			printlen = 40
		}
		count += 1
		fmt.Printf("%v\t[%v]\t%v:\t%v...\n", count, time.Now().Format("2006-01-02T15:04:05.999"), out.Url, printstr[:printlen])
	}
}

type Result struct {
	Error string
	Url   string
	Value string
}

func Fetch(url string, ch chan Result) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Error: err.Error(), Url: url}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	ch <- Result{Value: string(body), Url: url}
}
