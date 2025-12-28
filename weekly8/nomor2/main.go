package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// WebFetcher mensimulasikan fetch URL
func WebFetcher(url string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulasi delay 1–3 detik
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	ch <- "Fetched: " + url
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// daftar 5 URL dummy
	urls := []string{
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/3",
		"https://example.com/4",
		"https://example.com/5",
	}

	results := make(chan string)
	var wg sync.WaitGroup

	// goroutine penerima hasil
	go func() {
		for res := range results {
			log.Println(res)
		}
	}()

	// jalankan WebFetcher secara konkuren
	for _, url := range urls {
		wg.Add(1)
		go WebFetcher(url, results, &wg)
	}

	// tunggu semua fetch selesai
	wg.Wait()
	close(results)

	log.Println("Semua fetch selesai")
}

