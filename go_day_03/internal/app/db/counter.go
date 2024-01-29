package db

import (
	"bufio"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// GetCounts gets the number of all documents in the "place" index
func GetCounts() int {
	resp, _ := http.Get("http://127.0.0.1:9200/_cat/indices")
	defer resp.Body.Close()
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		res := strings.Split(sc.Text(), " ")
		if res[2] == "places" {
			c, err := strconv.Atoi(res[6])
			if err != nil {
				log.Fatal(err)
			}
			return c
		}
	}
	return 0
}
