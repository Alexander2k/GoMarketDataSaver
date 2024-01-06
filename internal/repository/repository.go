package repository

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/Alexander2k/CryptoBotGo/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error)
	SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error)
	SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error)
}

type Repository struct {
	IRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{IRepository: postgress.NewPostgresRepository(db)}
}
