package clickhouseRepo

import (
	"context"
	"github.com/Alexander2k/CryptoBotGo/internal/domain"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
	"log/slog"
)

type ClickHouseRepository struct {
	conn clickhouse.Conn
}

func NewClickHouseRepository(conn clickhouse.Conn) *ClickHouseRepository {
	return &ClickHouseRepository{conn: conn}
}

func (c *ClickHouseRepository) SaveSpotTicker(ctx context.Context, spot *domain.BybitTickersSpot) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClickHouseRepository) SavePerpetualTicker(ctx context.Context, perp *domain.BybitTickersPerp) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClickHouseRepository) SaveTrade(ctx context.Context, e *domain.Event, trade *domain.BybitTrade) (int64, error) {
	//TODO implement me
	panic("implement me")
}
func (c *ClickHouseRepository) SaveHeatMap(ctx context.Context, prices *domain.MeanPrices) error {

	for pr, qty := range prices.Prices {

		newUUID, _ := uuid.NewUUID()

		_, err := c.conn.Query(ctx, clkInsertHeatMap, newUUID,
			prices.Timestamp,
			prices.Market,
			prices.Ticker,
			pr,
			qty)

		if err != nil {
			slog.Error("Error SaveHeatMap", err)
			return err

		} else {
			slog.Info("Message saved")
		}

		//log.Printf("%v,%v,%v,%v,%v\n", prices.Timestamp,
		//	prices.Market,
		//	prices.Ticker,
		//	pr,
		//	qty)

	}

	return nil
}
