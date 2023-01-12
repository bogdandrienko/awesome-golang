package main

import (
	"encoding/json"
	"fmt"
	"golang-study/1_base/coincap"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return l.next.RoundTrip(req)
}

func http1() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("REDIRECTED")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 10,
	}

	const url1 string = "http://jsonplaceholder.typicode.com/todos/1/"
	const url2 string = "https://jsonplaceholder.typicode.com/todos/1/"
	resp, err := client.Get(url2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	defer resp.Body.Close()
}

func http2() {
	resp, err := http.DefaultClient.Get("https://api.coincap.io/v2/assets")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode, string(body))

	var r assetsResponse
	if json.Unmarshal(body, &r) != nil {
		log.Fatal(err)
	}

	for _, asset := range r.Assets {
		fmt.Println(asset.Info())
	}

}

type assetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupplySupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24h string `json:"volumeUsd24Hr"`
	PriceUSD     string `json:"PriceUsd"`
}

func (d assetData) Info() string {
	return fmt.Sprintf("[Id] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}

type assetsResponse struct {
	Assets    []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func http3() {
	resp, err := http.DefaultClient.Get("https://api.coincap.io/v2/assets/bitcoin")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode, string(body))

	var r assetResponse
	if json.Unmarshal(body, &r) != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Asset.Info())

}

type assetResponse struct {
	Asset     assetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

func http4() {
	//TODO array
	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	//TODO one
	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin.Info())
}

func main() {
	//http1()
	//http2()
	//http3()
	http4()
}
