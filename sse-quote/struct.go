package ssequote

/**
CLOUD SSE
*/

type Quote struct {
	Symbol                 string  `json:"symbol,omitempty"`
	CompanyName            string  `json:"companyName,omitempty"`
	PrimaryExchange        string  `json:"primaryExchange,omitempty"`
	CalculationPrice       string  `json:"calculationPrice,omitempty"`
	Open                   float64 `json:"open,omitempty"`
	OpenTime               int64   `json:"openTime,omitempty"`
	OpenSource             string  `json:"openSource,omitempty"`
	Close                  float64 `json:"close,omitempty"`
	CloseTime              int64   `json:"closeTime,omitempty"`
	CloseSource            string  `json:"closeSource,omitempty"`
	High                   float64 `json:"high,omitempty"`
	HighTime               int64   `json:"highTime,omitempty"`
	HighSource             string  `json:"highSource,omitempty"`
	Low                    float64 `json:"low,omitempty"`
	LowTime                int64   `json:"lowTime,omitempty"`
	LowSource              string  `json:"lowSource,omitempty"`
	LatestPrice            float64 `json:"latestPrice,omitempty"`
	LatestSource           string  `json:"latestSource,omitempty"`
	LatestTime             string  `json:"latestTime,omitempty"`
	LatestUpdate           int64   `json:"latestUpdate,omitempty"`
	LatestVolume           int     `json:"latestVolume,omitempty"`
	IexRealtimePrice       float64 `json:"iexRealtimePrice,omitempty"`
	IexRealtimeSize        int     `json:"iexRealtimeSize,omitempty"`
	IexLastUpdated         int64   `json:"iexLastUpdated,omitempty"`
	DelayedPrice           float64 `json:"delayedPrice,omitempty"`
	DelayedPriceTime       int64   `json:"delayedPriceTime,omitempty"`
	OddLotDelayedPrice     float64 `json:"oddLotDelayedPrice,omitempty"`
	OddLotDelayedPriceTime uint64  `json:"oddLotDelayedPriceTime,omitempty"`
	ExtendedPrice          float64 `json:"extendedPrice,omitempty"`
	ExtendedChange         float64 `json:"extendedChange,omitempty"`
	ExtendedChangePercent  float64 `json:"extendedChangePercent,omitempty"`
	ExtendedPriceTime      int64   `json:"extendedPriceTime,omitempty"`
	PreviousClose          float64 `json:"previousClose,omitempty"`
	PreviousVolume         int     `json:"previousVolume,omitempty"`
	Change                 float64 `json:"change,omitempty"`
	ChangePercent          float64 `json:"changePercent,omitempty"`
	Volume                 int     `json:"volume,omitempty"`
	IexMarketPercent       float64 `json:"iexMarketPercent,omitempty"`
	IexVolume              int     `json:"iexVolume,omitempty"`
	AvgTotalVolume         int     `json:"avgTotalVolume,omitempty"`
	IexBidPrice            float64 `json:"iexBidPrice,omitempty"`
	IexBidSize             int     `json:"iexBidSize,omitempty"`
	IexAskPrice            float64 `json:"iexAskPrice,omitempty"`
	IexAskSize             int     `json:"iexAskSize,omitempty"`
	MarketCap              int64   `json:"marketCap,omitempty"`
	PeRatio                float64 `json:"peRatio,omitempty"`
	Week52High             float64 `json:"week52High,omitempty"`
	Week52Low              float64 `json:"week52Low,omitempty"`
	YtdChange              float64 `json:"ytdChange,omitempty"`
	LastTradeTime          int64   `json:"lastTradeTime,omitempty"`
}

/**
SOCKET REQUEST
*/
type Request struct {
	Symbols string `json:"symbols"`
}

type Response struct {
	Success    bool        `json:"success"`
	Result     *ResultType `json:"result"`
	Error      *ErrorType  `json:"error"`
	ServerTime int64       `json:"serverTime"`
}

type ResultType struct {
	Quote *[]Quote `json:"data,omitempty"`
}

type ErrorType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
