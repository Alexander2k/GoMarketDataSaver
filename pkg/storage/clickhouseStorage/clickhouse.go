package clickhouseStorage

import (
	"context"
	"fmt"
	"github.com/Alexander2k/CryptoBotGo/config"
	"github.com/ClickHouse/clickhouse-go/v2"

	"time"
)

type ClickHouseDB struct {
	DB clickhouse.Conn
}

func NewClickHouseDB(cfg *config.Config) (*ClickHouseDB, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ClickDB.Host + ":" + cfg.ClickDB.Port},
		Auth: clickhouse.Auth{
			Database: cfg.ClickDB.Database,
			Username: cfg.ClickDB.Username,
			Password: cfg.ClickDB.Password,
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": cfg.ClickDB.MaxExecutionTime,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Minute * 1,
		MaxOpenConns:         cfg.ClickDB.MaxOpenConns,
		MaxIdleConns:         cfg.ClickDB.MaxIdleConns,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      cfg.ClickDB.BlockBufferSize,
		MaxCompressionBuffer: cfg.ClickDB.MaxCompressionBuffer,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "crypto-clickhouse", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &ClickHouseDB{DB: conn}, err
}
