package multiplex

import (
	"sync"
)

func Multiplex(chs ...<-chan interface{}) chan interface{} {
	result := make(chan interface{})

	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for el := range c {
				result <- el
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
