package domain

type (
	BybitKline struct {
		Topic string `json:"topic"`
		Data  []struct {
			Start     int64  `json:"start"`
			End       int64  `json:"end"`
			Interval  string `json:"interval"`
			Open      string `json:"open"`
			Close     string `json:"close"`
			High      string `json:"high"`
			Low       string `json:"low"`
			Volume    string `json:"volume"`
			Turnover  string `json:"turnover"`
			Confirm   bool   `json:"confirm"`
			Timestamp int64  `json:"timestamp"`
		} `json:"data"`
		Ts   int64  `json:"ts"`
		Type string `json:"type"`
	}

	BybitTrade struct {
		Topic string `json:"topic"`
		Type  string `json:"type"`
		Ts    int64  `json:"ts"`
		Data  []struct {
			T          int64  `json:"T"`
			Symbol     string `json:"s"`
			Side       string `json:"S"`
			TradeSize  string `json:"v"`
			TradePrice string `json:"p"`
			Direction  string `json:"L"`
			TradeId    string `json:"i"`
			BlockTrade bool   `json:"BT"`
		} `json:"data"`
	}

	BybitTickersPerp struct {
		Topic string `json:"topic"`
		Type  string `json:"type"`
		Data  struct {
			Symbol            string `json:"symbol"`
			TickDirection     string `json:"tickDirection"`
			Price24HPcnt      string `json:"price24hPcnt"`
			LastPrice         string `json:"lastPrice"`
			PrevPrice24H      string `json:"prevPrice24h"`
			HighPrice24H      string `json:"highPrice24h"`
			LowPrice24H       string `json:"lowPrice24h"`
			PrevPrice1H       string `json:"prevPrice1h"`
			MarkPrice         string `json:"markPrice"`
			IndexPrice        string `json:"indexPrice"`
			OpenInterest      string `json:"openInterest"`
			OpenInterestValue string `json:"openInterestValue"`
			Turnover24H       string `json:"turnover24h"`
			Volume24H         string `json:"volume24h"`
			NextFundingTime   string `json:"nextFundingTime"`
			FundingRate       string `json:"fundingRate"`
			Bid1Price         string `json:"bid1Price"`
			Bid1Size          string `json:"bid1Size"`
			Ask1Price         string `json:"ask1Price"`
			Ask1Size          string `json:"ask1Size"`
		} `json:"data"`
		Cs int64 `json:"cs"`
		Ts int64 `json:"ts"`
	}

	BybitTickersSpot struct {
		Topic string `json:"topic"`
		Ts    int64  `json:"ts"`
		Type  string `json:"type"`
		Cs    int64  `json:"cs"`
		Data  struct {
			Symbol        string `json:"symbol"`
			LastPrice     string `json:"lastPrice"`
			HighPrice24H  string `json:"highPrice24h"`
			LowPrice24H   string `json:"lowPrice24h"`
			PrevPrice24H  string `json:"prevPrice24h"`
			Volume24H     string `json:"volume24h"`
			Turnover24H   string `json:"turnover24h"`
			Price24HPcnt  string `json:"price24hPcnt"`
			UsdIndexPrice string `json:"usdIndexPrice"`
		} `json:"data"`
	}

	BybitBookTickerSpot struct {
		Topic        string `json:"topic"`
		TimeSendData int64  `json:"ts"`
		Type         string `json:"type"`
		Data         struct {
			TradingPair  string `json:"s"`
			BestBidPrice string `json:"bp"`
			BidQty       string `json:"bq"`
			BestAskPrice string `json:"ap"`
			AskQty       string `json:"aq"`
			TimeSysData  int64  `json:"t"`
		} `json:"data"`
	}

	ConnectionPropertyByBit struct {
		ReqId string   `json:"req_id,omitempty"`
		Op    string   `json:"op,omitempty"`
		Args  []string `json:"args,omitempty"`
	}

	PingMessage struct {
		ReqId string `json:"req_id,omitempty"`
		Op    string `json:"op"`
	}
)
