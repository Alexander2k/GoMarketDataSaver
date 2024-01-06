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
	sync.RWMutex
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
	meanPrices := NewMeanPrices()
	tmp := make(map[string]float64)
	b.mu.Lock()

	for k, _ := range b.Prices {
		prices := b.Prices[k]
		length := len(prices)

		sum := 0.0
		for _, p := range prices {
			float, err := strconv.ParseFloat(p, 64)
			if err != nil {
				return nil
			}
			sum += float
		}
		mean := sum / float64(length)
		tmp[k] = mean
	}

	b.ClearDataUnsafe()
	meanPrices.Prices = tmp
	meanPrices.Timestamp = time.Now().Unix()

	b.mu.Unlock()

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
