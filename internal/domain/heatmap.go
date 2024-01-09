package domain

import (
	"strconv"
	"sync"
	"time"
)

type BookStorage struct {
	Market string              `json:"market,omitempty"`
	Ticker string              `json:"ticker,omitempty"`
	Prices map[string][]string `json:"tickers,omitempty"`
	mu     sync.RWMutex
}

type MeanPrices struct {
	Market    string             `json:"market,omitempty"`
	Ticker    string             `json:"ticker,omitempty"`
	Timestamp int64              `json:"timestamp,omitempty"`
	Prices    map[string]float64 `json:"tickers,omitempty"`
}

func NewMeanPrices() *MeanPrices {
	return &MeanPrices{Prices: make(map[string]float64)}
}

func NewBookStorage() *BookStorage {
	return &BookStorage{Prices: make(map[string][]string)}
}

func (b *BookStorage) StorePrice(prices []string) {
	b.mu.Lock()
	b.Prices[prices[0]] = append(b.Prices[prices[0]], prices[1])
	b.mu.Unlock()
}

func (b *BookStorage) CalculateMeanPrice() *MeanPrices {

	b.mu.RLock()
	prices := make(map[string][]string, len(b.Prices))
	for k, v := range b.Prices {
		prices[k] = v
	}
	b.ClearDataUnsafe()
	b.mu.RUnlock()

	meanPrices := NewMeanPrices()
	tmp := make(map[string]float64)

	for k, _ := range prices {
		pricesQty := prices[k]
		length := len(pricesQty)

		sum := 0.0
		for _, p := range pricesQty {
			float, err := strconv.ParseFloat(p, 64)
			if err != nil {
				return nil
			}
			sum += float
		}
		mean := sum / float64(length)
		tmp[k] = mean
	}

	meanPrices.Prices = tmp
	meanPrices.Timestamp = time.Now().Unix()

	return meanPrices
}

func (b *BookStorage) ClearDataUnsafe() {
	for k := range b.Prices {
		delete(b.Prices, k)
	}

}
func (b *BookStorage) ClearDataSafe() {
	b.mu.Lock()
	for k := range b.Prices {
		delete(b.Prices, k)
	}
	b.mu.Unlock()
}
