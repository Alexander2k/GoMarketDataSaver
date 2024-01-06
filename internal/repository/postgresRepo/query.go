package postgresRepo

var (
	sqlSavePerpTicker = `insert into tickers_perp(
                         topic,
                         type,
                         symbol,
                         tickerdirection,
                         price24hpcnt,
                         lastprice,
                         prevprice24h,
                         highprice24h,
                         lowprice24h,
                         prevprice1h,
                         markprice,
                         indexprice,
                         openinterest,
                         openinterestvalue,
                         turnover24h,
                         volume24h,
                         nextfundingtime,
                         fundingrate,
                         bid1price,
                         bid1size,
                         ask1price,
                         ask1size,
                         cs,
                         ts) values ($1, $2, $3, $4, $5, $6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24) returning id`

	sqlSaveSpotTicker = `insert into tickers_spot(
                         topic,
                         ts,
                         type,
                         cs,
                         symbol,
                         lastprice,
                         highprice24h,
                         lowprice24h,
                         prevprice24h,
                         volume24h,
                         turnover24h,
                         price24hpcnt,
                         usdindexprice) values ($1, $2, $3, $4, $5, $6,$7,$8,$9,$10,$11,$12,$13) returning id`

	sqlSaveTrades = `insert into trade(
                  market,
                  topic,
                  type,
                  ts,
                  t,
                  symbol,
                  side,
                  trade_size,
                  trade_price,
                  direction,
                  tradeid,
                  blocktrade) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id`
)
