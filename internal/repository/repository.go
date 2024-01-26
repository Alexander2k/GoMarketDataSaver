package repository

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/Alexander2k/CryptoBotGo/internal/repository/postgresRepo"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
)

type PgRepository interface {
	SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error)
	SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error)
	SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error)
	SaveHeatMap(ctx context.Context, prices *domain.MeanPrices) error
}

type ClickHouseRepository interface {
	SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error)
	SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error)
	SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error)
	SaveHeatMap(ctx context.Context, prices *domain.MeanPrices) error
}

type Repository struct {
	PgRepository
	ClickHouseRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PgRepository: postgresRepo.NewPostgresRepository(db),
		//ClickHouseRepository: clickhouseRepo.NewClickHouseRepository(conn),
	}
}
