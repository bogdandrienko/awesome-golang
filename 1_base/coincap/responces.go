package coincap

import "fmt"

type Asset struct {
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

func (d Asset) Info() string {
	return fmt.Sprintf("[Id] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}
