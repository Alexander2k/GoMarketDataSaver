package main

import (
	"fmt"
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/Alexander2k/CryptoBotGo/internal/exchange"
	"github.com/Alexander2k/CryptoBotGo/internal/repository"
	"github.com/Alexander2k/CryptoBotGo/pkg/storage/clickhouseStorage"
	"github.com/Alexander2k/CryptoBotGo/pkg/storage/postgresStorage"
	"net/http"
	"time"
)

func Start() error {
	//config init
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
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
		return err
	}
	if err = db.Migrate(); err != nil {
		return err
	}

	clickHouseDB, err := clickhouseStorage.NewClickHouseDB(conf)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db.Db, clickHouseDB.DB)
	ex := exchange.NewExchange(repo)

	bybitPerp := ex.BybitConnectPerpetual(conf)

	go func() {
		for {
			orderBook, err := ex.CollectOrderBook(bybitPerp)
			if err != nil {
				return
			}
			fmt.Printf("%v\n", orderBook)
		}

	}()

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
