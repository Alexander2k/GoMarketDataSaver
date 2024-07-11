package main

import (
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/Alexander2k/CryptoBotGo/internal/exchange"
	"github.com/Alexander2k/CryptoBotGo/internal/metrics"
	"github.com/Alexander2k/CryptoBotGo/internal/repository"
	"github.com/Alexander2k/CryptoBotGo/pkg/storage/postgresStorage"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {

	go func() {
		err := metrics.Listen("localhost:8989")
		if err != nil {
			log.Printf("Error listening prometheus metrics: %v", err)
			return
		}
	}()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	conf, err := config.NewConfig()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	server := http.Server{
		Addr:              "localhost:8181",
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 60,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 60,
	}

	db, err := postgresStorage.NewPostgresDB(conf)
	if err != nil {
		logger.Error(err.Error())
	}
	if err = db.Migrate(); err != nil {
		logger.Info(err.Error())
		return
	}

	repo := repository.NewRepository(db.Db)
	ex := exchange.NewExchange(repo)

	bybitPerp := ex.BybitConnectPerpetual(conf)
	//bybitSpot := ex.BybitConnectSpot(conf)

	orderBookP, _, _, _, _ := ex.CollectData(bybitPerp)
	//_, _, _, _, _ := ex.CollectData(bybitSpot)

	//_ = ex.CollectTrades(tradesChanS, tradesChanP)
	//_ = ex.CollectTicker(tickerS, tickerP)
	//_ = ex.CollectCandle(candleS, candleP)'
	_ = ex.CollectOrderBook(orderBookP)
	//for {
	//	select {
	//	case x := <-orderBookP:
	//		log.Println(x.String())
	//	case y := <-orderBookS:
	//		log.Println(y.String())
	//	}
	//
	//}

	if err := server.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		return
	}

}
