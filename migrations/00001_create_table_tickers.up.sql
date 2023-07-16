create table if not exists tickers_perp
(
    id                bigserial primary key,
    topic             varchar default null,
    type              varchar default null,
    symbol            varchar default null,
    tickerDirection   varchar default null,
    price24hPcnt      varchar default null,
    lastPrice         varchar default null,
    prevPrice24h      varchar default null,
    highPrice24h      varchar default null,
    lowPrice24h       varchar default null,
    prevPrice1h       varchar default null,
    markPrice         varchar default null,
    indexPrice        varchar default null,
    openInterest      varchar default null,
    openInterestValue varchar default null,
    turnover24h       varchar default null,
    volume24h         varchar default null,
    nextFundingTime   varchar default null,
    fundingRate       varchar default null,
    Bid1Price         varchar default null,
    Bid1Size          varchar default null,
    Ask1Price         varchar default null,
    Ask1Size          varchar default null,
    cs                varchar  default null,
    ts                varchar  default null
);

create table if not exists tickers_spot
(
    id            bigserial primary key,
    topic         varchar,
    ts            varchar,
    type          varchar,
    cs            varchar,
    symbol        varchar,
    lastPrice     varchar,
    highPrice24h  varchar,
    lowPrice24h   varchar,
    prevPrice24h  varchar,
    volume24h     varchar,
    Turnover24H   varchar,
    Price24HPcnt  varchar,
    UsdIndexPrice varchar
);

create table if not exists trade
(
    id          bigserial primary key,
    market      varchar,
    topic       varchar,
    type        varchar,
    ts          varchar,
    t           varchar,
    symbol      varchar,
    side        varchar,
    trade_size  varchar,
    trade_price varchar,
    direction   varchar null,
    tradeid     varchar,
    blocktrade  bool
);

create index tickers_spot_symbol_index on tickers_spot(symbol);
create index tickers_spot_topic_index on tickers_spot(topic);
create index tickers_spot_ts_index on tickers_spot(ts);

create index tickers_perp_symbol_index on tickers_perp(symbol);
create index tickers_perp_topic_index on tickers_perp(topic);
create index tickers_perp_ts_index on tickers_perp(ts);

create index trade_symbol_index on trade(symbol);
create index trade_market_index on trade(market);
create index trade_side_index on trade(side);
create index trade_topic_index on trade(topic);
create index trade_ts_index on trade(ts);
