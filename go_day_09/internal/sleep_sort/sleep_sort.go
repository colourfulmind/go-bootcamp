package sleep_sort

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func SleepSort(arr []int) chan int {
	ch := make(chan int)

	if len(arr) == 0 {
		close(ch)
		return ch
	}

	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			time.Sleep(time.Duration(value) * time.Second)
			ch <- value
		}(v)

		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			select {
			case sgl := <-exit:
				fmt.Printf("gracefully stopped with signal %v\n", sgl)
				close(ch)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
