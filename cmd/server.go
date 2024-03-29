package main

import (
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/Alexander2k/CryptoBotGo/internal/exchange"
	"github.com/Alexander2k/CryptoBotGo/internal/repository"
	"github.com/Alexander2k/CryptoBotGo/pkg/storage/postgres"
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

	db, err := postgres.NewPostgresDB(conf)
	if err != nil {
		return err
	}
	if err = db.Migrate(); err != nil {
		return err
	}

	repo := repository.NewRepository(db.Db)

	ex := exchange.NewExchange(repo)

	bybitPerp := ex.BybitConnectPerpetual(conf)
	bybitSpot := ex.BybitConnectSpot(conf)

	c := ex.ConveyorExchange(bybitPerp, bybitSpot)

	ex.StoreData(c)

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
