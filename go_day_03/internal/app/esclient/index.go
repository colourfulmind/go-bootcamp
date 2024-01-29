package esclient

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// CheckIndex checks if the index already exists
func CheckIndex(es *elasticsearch.Client) {
	ctx := context.Background()
	resp, err := esapi.IndicesExistsRequest{
		Index: []string{index},
	}.Do(ctx, es)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode == http.StatusNotFound {
		CreateRequest(es, ctx)
	}
}

// CreateRequest creates a request using `esapi`
func CreateRequest(es *elasticsearch.Client, ctx context.Context) {
	places := LoadData()
	mapping := LoadJSON()
	resp, err := esapi.IndicesCreateRequest{
		Index:  index,
		Body:   strings.NewReader("{\"mappings\": " + mapping + "}"),
		Pretty: true,
		Human:  true,
	}.Do(ctx, es)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	CreateIndex(es, ctx, places)
}

// LoadData loads data from the database
func LoadData() []Place {
	file, err := os.Open(databasePath)
	defer file.Close()
	rd := csv.NewReader(file)
	rd.Comma = del
	data, err := rd.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return CreatePlacesList(data)
}

// LoadJSON loads a JSON schema
func LoadJSON() string {
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// CreateIndex creates an index named "places"
func CreateIndex(es *elasticsearch.Client, ctx context.Context, places []Place) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  index,
		Client: es,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer bi.Close(ctx)
	for i, p := range places {
		data, err := json.Marshal(p)
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: strconv.Itoa(i + 1),
			Body:       bytes.NewReader(data),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
