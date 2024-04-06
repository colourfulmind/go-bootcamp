package crawler

import (
	"context"
	"io"
	"net/http"
	"sync"
)

const MaxGoroutines = 8

func CrawlWeb(ctx context.Context, urls <-chan string) chan *string {
	result := make(chan *string)
	ch := make(chan struct{}, MaxGoroutines)

	var wg sync.WaitGroup
	go func() {
		for url := range urls {
			select {
			case <-ctx.Done():
				break
			case ch <- struct{}{}:
				wg.Add(1)
				go func(u string) {
					defer wg.Done()
					body, err := GetRespBody(u)
					if err == nil {
						result <- &body
					}
					<-ch
				}(url)
			}
		}
		wg.Wait()
		close(result)
	}()

	return result
}

func GetRespBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
