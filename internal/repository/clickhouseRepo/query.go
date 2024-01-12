package clickhouseRepo

var (
	clkInsertHeatMap = `INSERT INTO buybit.heatmap (id,
                            timestamp,
                            market,
                            ticker,
                            price,
                            qty)
values (generateUUIDv4(),
        $2,
        $3,
        $4,
        $5,
        $6)`
)
