package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	urls := []string{"https://golang.org", "https://example.com", "https://nonexistent-site.xyz",
		"https://www.youtube.com", "https://www.meow.com/",
	}
	titles := make(map[string]string)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			client := http.Client{
				Timeout: 5 * time.Second,
			}
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("Error at making request:", err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error at getting resp:", err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Println("Error at getting doc from goquery:", err)
				return
			}

			title := doc.Find("title").First().Text()

			mu.Lock()
			titles[url] = title
			mu.Unlock()

		}(url)

	}

	wg.Wait()

	for key, value := range titles {
		fmt.Printf("URL: %v Title:%v\n", key, value)
	}

}
