package postgres

import (
	"fmt"
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/jmoiron/sqlx"
	"log"
)

type PostgresDB struct {
	Db *sqlx.DB
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PDB.Host,
		cfg.PDB.Port,
		cfg.PDB.Username,
		cfg.PDB.DBName,
		cfg.PDB.Password,
		cfg.PDB.SSLMode,
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

	return &PostgresDB{
		Db: db,
	}, nil

}
