package repository

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error) {

	result, err := r.db.ExecContext(ctx, sqlSaveSpotTicker,
		spot.Topic,
		spot.Ts,
		spot.Type,
		spot.Cs,
		spot.Data.Symbol,
		spot.Data.LastPrice,
		spot.Data.HighPrice24H,
		spot.Data.LowPrice24H,
		spot.Data.PrevPrice24H,
		spot.Data.Volume24H,
		spot.Data.Turnover24H,
		spot.Data.Price24HPcnt,
		spot.Data.UsdIndexPrice)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (r *Repository) SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error) {

	result, err := r.db.ExecContext(ctx, sqlSavePerpTicker,
		perp.Topic,
		perp.Type,
		perp.Data.Symbol,
		perp.Data.TickDirection,
		perp.Data.Price24HPcnt,
		perp.Data.LastPrice,
		perp.Data.PrevPrice24H,
		perp.Data.HighPrice24H,
		perp.Data.LowPrice24H,
		perp.Data.PrevPrice1H,
		perp.Data.MarkPrice,
		perp.Data.IndexPrice,
		perp.Data.OpenInterest,
		perp.Data.OpenInterestValue,
		perp.Data.Turnover24H,
		perp.Data.Volume24H,
		perp.Data.NextFundingTime,
		perp.Data.FundingRate,
		perp.Data.Bid1Price,
		perp.Data.Bid1Size,
		perp.Data.Ask1Price,
		perp.Data.Ask1Size,
		perp.Cs,
		perp.Ts)
	if err != nil {
		log.Printf("Error saving ticker: %v", err)
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error saving ticker no affected: %v", err)
		return 0, err
	}
	return affected, nil
}

func (r *Repository) SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error) {
	result, err := r.db.ExecContext(ctx, sqlSaveTrades,
		e.Market,
		trade.Topic,
		trade.Type,
		trade.Ts,
		trade.Data[0].T,
		trade.Data[0].Symbol,
		trade.Data[0].Side,
		trade.Data[0].TradeSize,
		trade.Data[0].TradePrice,
		trade.Data[0].Direction,
		trade.Data[0].TradeId,
		trade.Data[0].BlockTrade)
	if err != nil {
		log.Printf("Error saving trade: %v", err)
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error adding trade: %v", err)
		return 0, err
	}
	return affected, nil
}
