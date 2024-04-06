package output

import (
	"context"
	"fmt"
	"main/internal/crawler"
	"main/internal/multiplex"
	"main/internal/sleep_sort"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func SleepSortOutput() {
	arr := []int{6, 3, 4, 2, 5, 9, 8, 7, 0, 1}
	for el := range sleep_sort.SleepSort(arr) {
		fmt.Println(el)
	}
}

var urls = []string{
	"bufio@go1.22.2",
	"builtin@go1.22.2",
	"bytes@go1.22.2",
	"cmp@go1.22.2",
	"context@go1.22.2",
	"embed@go1.22.2",
	"errors@go1.22.2",
	"expvar@go1.22.2",
	"flag@go1.22.2",
	"fmt@go1.22.2",
}

func CrawlerOutput() {
	ch := make(chan string)
	go func() {
		for _, url := range urls {
			ch <- fmt.Sprintf("https://pkg.go.dev/%s", url)
		}
		close(ch)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sgl := <-stop
		fmt.Printf("gracefully stopped with signal %v\n", sgl)
		cancel()
	}()

	for el := range crawler.CrawlWeb(ctx, ch) {
		fmt.Println(el)
	}
}

const Total = 10

func MultiplexOutput() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		for i := 0; i < Total; i++ {
			ch1 <- 1 * i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < Total; i++ {
			ch2 <- "hello" + strconv.Itoa(i)
		}
		close(ch2)
	}()

	go func() {
		for i := 0.0; i < Total; i++ {
			ch3 <- 100.001 + i*i
		}
		close(ch3)
	}()

	res := multiplex.Multiplex(ch1, ch2, ch3)

	for el := range res {
		fmt.Println(el)
	}
}
