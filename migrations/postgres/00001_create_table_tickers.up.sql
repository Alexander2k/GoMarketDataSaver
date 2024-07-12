create table if not exists tickers_perp
(
    id                bigserial primary key,
    topic             varchar ,
    type              varchar ,
    symbol            varchar ,
    tickerDirection   varchar ,
    price24hPcnt      double precision,
    lastPrice         double precision,
    prevPrice24h      double precision,
    highPrice24h      double precision,
    lowPrice24h       double precision,
    prevPrice1h       double precision,
    markPrice         double precision,
    indexPrice        double precision,
    openInterest      double precision,
    openInterestValue double precision,
    turnover24h       double precision,
    volume24h         double precision,
    nextFundingTime   bigint,
    fundingRate       double precision,
    Bid1Price         double precision,
    Bid1Size          double precision,
    Ask1Price         double precision,
    Ask1Size          double precision,
    cs                bigint,
    ts                bigint
);

create table if not exists tickers_spot
(
    id            bigserial primary key,
    topic         varchar,
    ts            bigint,
    type          varchar,
    cs            bigint,
    symbol        varchar,
    lastPrice     double precision,
    highPrice24h  double precision,
    lowPrice24h   double precision,
    prevPrice24h  double precision,
    volume24h     double precision,
    Turnover24H   double precision,
    Price24HPcnt  double precision,
    UsdIndexPrice double precision
);


create table if not exists trades
(
    id          bigserial primary key,
    market      varchar,
    topic       varchar,
    type        varchar,
    ts          bigint,
    t           bigint,
    symbol      varchar,
    side        varchar,
    trade_size  double precision,
    trade_price double precision,
    direction   varchar null,
    tradeid     varchar,
    blocktrade  bool
);

create table if not exists orderbook
(
    id        bigserial primary key,
    timestamp timestamp,
    market    varchar,
    ticker    varchar,
    price     double precision,
    qty       double precision[]
);

create table if not exists candle_spot
(
    id          bigserial primary key,
    topic       varchar,
    startCandle bigint,
    endCandle   bigint,
    interval    integer,
    open        double precision,
    close       double precision,
    high        double precision,
    low         double precision,
    volume      double precision,
    turnover    double precision,
    confirm     boolean,
    timestamp   bigint,
    ts          bigint,
    type        varchar
);

create table if not exists candle_perp
(
    id          bigserial primary key,
    topic       varchar,
    startCandle bigint,
    endCandle   bigint,
    interval    integer,
    open        double precision,
    close       double precision,
    high        double precision,
    low         double precision,
    volume      double precision,
    turnover    double precision,
    confirm     boolean,
    timestamp   bigint,
    ts          bigint,
    type        varchar
);

