package postgres

import (
	"fmt"
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/jmoiron/sqlx"
	"log"
)

type DB struct {
	Db *sqlx.DB
}

func NewPostgresDB(cfg *config.Config) (*DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Username,
		cfg.Db.DBName,
		cfg.Db.Password,
		cfg.Db.SSLMode,
	))
	if err != nil {
		log.Printf("Error connecting: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging: %v", err)
		return nil, err
	}

	return &DB{
		Db: db,
	}, nil

}
