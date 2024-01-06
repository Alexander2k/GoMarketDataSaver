package repository

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/Alexander2k/CryptoBotGo/internal/repository/clickhouseRepo"
	"github.com/Alexander2k/CryptoBotGo/internal/repository/postgresRepo"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
)

type PgRepository interface {
	SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error)
	SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error)
	SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error)
}

type ClickHouseRepository interface {
	SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error)
	SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error)
	SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error)
}

type Repository struct {
	PgRepository
	ClickHouseRepository
}

func NewRepository(db *sqlx.DB, conn clickhouse.Conn) *Repository {
	return &Repository{
		PgRepository:         postgresRepo.NewPostgresRepository(db),
		ClickHouseRepository: clickhouseRepo.NewClickHouseRepository(conn),
	}
}
