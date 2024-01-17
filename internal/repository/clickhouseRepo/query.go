package clickhouseRepo

var (
	clkInsertHeatMap = `INSERT INTO buybit.heatmap (id,
                            timestamp,
                            market,
                            ticker,
                            price,
                            qtys) 
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6)`
)
