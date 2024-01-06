package clickhouse

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouseRepository struct {
	conn clickhouse.Conn
}

func NewClickHouseRepository(conn clickhouse.Conn) *ClickHouseRepository {
	return &ClickHouseRepository{conn: conn}
}

func (c ClickHouseRepository) SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClickHouseRepository) SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClickHouseRepository) SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error) {
	//TODO implement me
	panic("implement me")
}
