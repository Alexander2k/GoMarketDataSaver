package exchange

import (
	"encoding/json"
	"fmt"
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"log"
	"sync"
	"time"
)

func (e *Exchange) BybitConnectPerpetual(c *config.Config) chan *domain.Event {
	//dur := time.Minute * 1
	ticker := time.NewTicker(time.Minute * 5)
	messageChan := make(chan []byte, 1)
	dataChan := make(chan *domain.Event, 5)
	var wg = sync.WaitGroup{}
	wg.Add(2)

	dial, _, err := e.websocket.Dial(c.Bybit.UrlPerpetual, nil)
	if err != nil {
		log.Default().Printf("Err: %v", err)
	}

	subscribe, err := json.Marshal(&domain.ConnectionPropertyByBit{
		Op:   "subscribe",
		Args: c.Bybit.PerpetualPairs,
	})

	ping, err := json.Marshal(&domain.PingMessage{
		ReqId: "999999",
		Op:    "ping",
	})

	err = dial.WriteMessage(1, subscribe)
	if err != nil {
		return nil
	}
	go func() {
		defer wg.Done()
		for {
			_, m, err := dial.ReadMessage()
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			messageChan <- m

		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				err := dial.WriteMessage(1, ping)
				if err != nil {
					log.Printf("Error: %v", err)
				} else {
					log.Printf("Sending ping")
				}
			case mc := <-messageChan:
				dataChan <- &domain.Event{
					Market: "Perpetual",
					Event:  mc,
				}
			}
		}

	}()
	go func() {
		wg.Wait()
	}()

	return dataChan

}

func (e *Exchange) BybitConnectSpot(c *config.Config) chan *domain.Event {

	ticker := time.NewTicker(time.Minute * 5)
	messageChan := make(chan []byte, 1)
	dataChan := make(chan *domain.Event, 5)
	var wg = sync.WaitGroup{}
	wg.Add(2)

	dial, _, err := e.websocket.Dial(c.Bybit.UrlSpot, nil)
	if err != nil {
		log.Printf("Error dial connection: %v", err)
	}

	subscribe, err := json.Marshal(&domain.ConnectionPropertyByBit{
		Op:   "subscribe",
		Args: c.Bybit.SpotPairs,
	})

	ping, err := json.Marshal(&domain.PingMessage{
		ReqId: "999999",
		Op:    "ping",
	})

	err = dial.WriteMessage(1, subscribe)
	if err != nil {
		log.Printf("Error subscription: %v", err)
	}
	go func() {
		defer wg.Done()
		for {

			_, m, err := dial.ReadMessage()
			if err != nil {
				log.Printf("Error read messahege: %v", err)
			}
			messageChan <- m
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				err := dial.WriteMessage(1, ping)
				if err != nil {
					log.Printf("Error: %v", err)
				}
			case m := <-messageChan:
				dataChan <- &domain.Event{
					Market: "Spot",
					Event:  m,
				}
			}

		}
	}()

	go func() { wg.Wait() }()
	return dataChan

}
