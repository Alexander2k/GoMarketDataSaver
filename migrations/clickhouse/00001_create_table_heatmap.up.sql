CREATE DATABASE if not exists buybit;

CREATE TABLE if not exists buybit.heatmap
(
    id        UUID,
    timestamp DATETIME,
    market    String,
    ticker    String,
    price     String,
    qty       Float64

)
    ENGINE = MergeTree()
        PRIMARY KEY (id, timestamp, ticker, price);