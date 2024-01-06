package repository

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	clk "github.com/Alexander2k/CryptoBotGo/internal/repository/clickhouse"
	"github.com/Alexander2k/CryptoBotGo/internal/repository/postgress"
	"github.com/ClickHouse/clickhouse-go/v2"
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
		PgRepository:         postgress.NewPostgresRepository(db),
		ClickHouseRepository: clk.NewClickHouseRepository(conn),
	}
}
