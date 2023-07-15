package exchange

import (
	"context"
	"encoding/json"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/Alexander2k/CryptoBotGo/internal/repository"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"sync"
)

type Exchange struct {
	websocket *websocket.Dialer
	Repo      *repository.Repository
}

func NewExchange(repo *repository.Repository) *Exchange {
	return &Exchange{Repo: repo}
}

func (e *Exchange) ConveyorExchange(channels ...<-chan *domain.Event) chan *domain.Event {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	dataChan := make(chan *domain.Event, 20)
	for _, channel := range channels {
		ch := channel
		go func() {
			defer wg.Done()
			for data := range ch {
				dataChan <- data
			}
		}()
	}
	go func() {
		wg.Wait()
	}()

	return dataChan
}

func (e *Exchange) StoreData(channels ...<-chan *domain.Event) {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, channel := range channels {
		ch := channel
		go func() {
			defer wg.Done()
			for data := range ch {
				switch data.Market {
				case "Perpetual":
					if strings.Contains(string(data.Event), "tickers") {
						var ticker domain.BybitTickersPerp
						err := json.Unmarshal(data.Event, &ticker)
						if err != nil {
							log.Printf("Error Unmarshal Perpetual: %v, %v", err, string(data.Event))
						}
						_, err = e.Repo.SavePerpetualTicker(context.Background(), &ticker)
						if err != nil {
							log.Printf("Error saving perpetual ticker: %v - %v", err, string(data.Event))
						}
					}

					if strings.Contains(string(data.Event), "publicTrade") {
						var trade domain.BybitTrade
						err := json.Unmarshal(data.Event, &trade)
						if err != nil {
							log.Printf("Error Unmarshal Perpetual publicTrade data: %v, %v", err, string(data.Event))
						}

						_, err = e.Repo.SaveTrade(context.Background(), data, &trade)
						if err != nil {
							log.Printf("Error saving Perpetual trade data: %v - %v", err, string(data.Event))
						}
					}

				case "Spot":
					if strings.Contains(string(data.Event), "tickers") {
						var ticker domain.BybitTickersSpot
						err := json.Unmarshal(data.Event, &ticker)
						if err != nil {
							log.Printf("Error Unmarshal Spot: %v%v", err, string(data.Event))
						}

						_, err = e.Repo.SaveSpotTicker(context.Background(), &ticker)
						if err != nil {
							log.Printf("Error saving spot ticker: %v - %v", err, string(data.Event))
						}
					}

					if strings.Contains(string(data.Event), "publicTrade") {
						var trade domain.BybitTrade
						err := json.Unmarshal(data.Event, &trade)
						if err != nil {
							log.Printf("Error Unmarshal publicTrade data: %v - %v", err, string(data.Event))
						}
						_, err = e.Repo.SaveTrade(context.Background(), data, &trade)
						if err != nil {
							log.Printf("Error saving Spot trade data: %v - %v", err, string(data.Event))
						}
					}

				}

			}
		}()

		go func() {
			wg.Wait()
		}()
	}

}
