package cliclhouse

import (
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/jmoiron/sqlx"
)

type ClickHouseDB struct {
	db *sqlx.DB
}

func NewClickHouseDB(cfg *config.Config) (*ClickHouseDB, error) {

	return &ClickHouseDB{}, nil
}
