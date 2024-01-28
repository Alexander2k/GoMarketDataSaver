package exchange

import (
	"context"
	"encoding/json"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/Alexander2k/CryptoBotGo/internal/repository"
	"github.com/gorilla/websocket"
	"log"
	"log/slog"
	"strings"
	"sync"
	"time"
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
	dataChan := make(chan *domain.Event, 5)
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
	storage := domain.NewBookStorage()
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
						_, err = e.Repo.PgRepository.SavePerpetualTicker(context.Background(), &ticker)
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

						_, err = e.Repo.PgRepository.SaveTrade(context.Background(), data, &trade)
						if err != nil {
							log.Printf("Error saving Perpetual trade data: %v - %v", err, string(data.Event))
						}
					}

					if strings.Contains(string(data.Event), "orderbook") {
						var orderBook domain.BybitOrderBook
						err := json.Unmarshal(data.Event, &orderBook)
						if err != nil {
							log.Printf("Error Unmarshal Perpetual")
						}

						storage.Market = data.Market

						if orderBook.Type == "snapshot" {
							for _, a := range orderBook.Data.Asks {
								storage.StorePrice(a)
							}

							for _, b := range orderBook.Data.Bids {
								storage.StorePrice(b)
							}

						}

						if orderBook.Type == "delta" {
							for _, a := range orderBook.Data.Asks {
								storage.StorePrice(a)
							}

							for _, b := range orderBook.Data.Bids {
								storage.StorePrice(b)
							}

						}
					}

				case "Spot":
					if strings.Contains(string(data.Event), "tickers") {
						var ticker domain.BybitTickersSpot
						err := json.Unmarshal(data.Event, &ticker)
						if err != nil {
							log.Printf("Error Unmarshal Spot: %v%v", err, string(data.Event))
						}

						_, err = e.Repo.PgRepository.SaveSpotTicker(context.Background(), &ticker)
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
						_, err = e.Repo.PgRepository.SaveTrade(context.Background(), data, &trade)
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

func (e *Exchange) CollectOB(channel chan *domain.Event) (*domain.MeanPrices, error) {
	var orderBook domain.BybitOrderBook
	storage := domain.NewBookStorage()

	ticker := time.NewTicker(60000 * time.Millisecond)

	for {
		select {
		case x := <-channel:
			{
				if strings.Contains(string(x.Event), "orderbook") {
					err := json.Unmarshal(x.Event, &orderBook)
					if err != nil {
						return nil, err
					}

					storage.Market = x.Market

					if orderBook.Type == "snapshot" {
						for _, a := range orderBook.Data.Asks {
							storage.StorePrice(a)
						}

						for _, b := range orderBook.Data.Bids {
							storage.StorePrice(b)
						}

					}

					if orderBook.Type == "delta" {
						for _, a := range orderBook.Data.Asks {
							storage.StorePrice(a)
						}

						for _, b := range orderBook.Data.Bids {
							storage.StorePrice(b)
						}

					}
				}
			}
		case x := <-ticker.C:
			slog.Info("Event", x)
			data := storage.CalculateMeanPrice()
			topic := strings.Split(orderBook.Topic, ".")
			data.Ticker = topic[2]
			data.Market = storage.Market
			return data, nil

		}
	}

}

func (e *Exchange) CollectData(in chan *domain.Event) (chan *domain.Event, chan *domain.Event, chan *domain.Event, chan *domain.Event, chan *domain.Event) {
	var wg sync.WaitGroup
	orderBookChannel := make(chan *domain.Event)
	candleChanel := make(chan *domain.Event)
	tradesChanel := make(chan *domain.Event)
	tickerChanel := make(chan *domain.Event)
	liquidChanel := make(chan *domain.Event)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case x := <-in:
				if strings.Contains(string(x.Event), "orderbook") {
					orderBookChannel <- x
				}

				if strings.Contains(string(x.Event), "kline") {
					candleChanel <- x
				}

				if strings.Contains(string(x.Event), "publicTrade") {
					tradesChanel <- x
				}
				if strings.Contains(string(x.Event), "tickers") {
					tickerChanel <- x
				}
				if strings.Contains(string(x.Event), "liquidation") {
					liquidChanel <- x
				}

			}
		}
	}()

	go func() {
		wg.Wait()
	}()

	return orderBookChannel, candleChanel, tradesChanel, tickerChanel, liquidChanel

}

func (e *Exchange) CollectOrderBook(in chan *domain.Event) error {
	var orderBook domain.BybitOrderBook
	var wg sync.WaitGroup
	storage := domain.NewBookStorage()

	ticker := time.NewTicker(60000 * time.Millisecond)

	wg.Add(1)
	go func() {
		for {
			defer wg.Done()
			select {
			case x := <-in:
				{
					err := json.Unmarshal(x.Event, &orderBook)
					if err != nil {
						return
					}

					storage.Market = x.Market

					if orderBook.Type == "snapshot" {
						for _, a := range orderBook.Data.Asks {
							storage.StorePrice(a)
						}

						for _, b := range orderBook.Data.Bids {
							storage.StorePrice(b)
						}

					}

					if orderBook.Type == "delta" {
						for _, a := range orderBook.Data.Asks {
							storage.StorePrice(a)
						}

						for _, b := range orderBook.Data.Bids {
							storage.StorePrice(b)
						}

					}

				}
			case <-ticker.C:
				data := storage.CalculateMeanPrice()
				topic := strings.Split(orderBook.Topic, ".")
				data.Ticker = topic[2]
				data.Market = storage.Market

				log.Println("Length data", data.Prices)
				err := e.Repo.PgRepository.SaveHeatMap(context.Background(), data)
				if err != nil {
					return
				}

			}
		}
	}()

	go func() {
		wg.Wait()
	}()

	return nil

}

func (e *Exchange) CollectCandle(in chan *domain.Event) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case event := <-in:
				if strings.Contains(string(event.Event), "true") {
					log.Println("Event kline", string(event.Event))
					log.Println("Save kline")
				}

			}
		}

	}()

	go func() {
		wg.Wait()
	}()
	return nil

}

func (e *Exchange) CollectTrades(in chan *domain.Event) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case event := <-in:
				log.Println("Event Trades", string(event.Event))
				log.Println("Save trade")
			}
		}

	}()

	go func() {
		wg.Wait()
	}()

	return nil

}

func (e *Exchange) CollectTicker(in chan *domain.Event) error {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case event := <-in:
				log.Println("Event Ticker", string(event.Event))
				log.Println("Save Tick")
			}
		}
	}()

	go func() {
		wg.Wait()
	}()

	return nil
}

func (e *Exchange) CollectLiquidation(in chan *domain.Event) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case event := <-in:
				log.Println("Event Liquidation", string(event.Event))
				log.Println("Save Liquidation")
			}
		}
	}()

	go func() {
		wg.Wait()
	}()

	return nil
}
