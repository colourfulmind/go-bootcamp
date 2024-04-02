package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CertFile = "cert/client/cert.pem"
	KeyFile  = "cert/client/key.pem"
	MiniCA   = "cert/minica.pem"
)

var (
	Type  string
	Count int
	Money int
)

type Request struct {
	Type  string `json:"candyType"`
	Count int    `json:"candyCount"`
	Money int    `json:"money"`
}

func init() {
	flag.StringVar(&Type, "k", "", "type of candy")
	flag.IntVar(&Count, "c", 0, "candy amount")
	flag.IntVar(&Money, "m", 0, "money amount")
	flag.Parse()
}

func main() {
	CACert, err := GetTLSConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: CACert,
		},
	}

	fmt.Println(MakeRequest(client))
}

func MakeRequest(client *http.Client) string {
	req, _ := json.Marshal(Request{
		Type:  Type,
		Count: Count,
		Money: Money,
	})

	data, err := client.Post("https://candy.tld:3333/buy_candy", "application/json", strings.NewReader(string(req)))
	defer data.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp)
}

func GetTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		return nil, err
	}

	CACert, err := os.ReadFile(MiniCA)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(CACert)

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	return config, nil
}
