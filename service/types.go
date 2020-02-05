package service

//Binance is the base type for the binance exchange
type Binance struct {
	BaseURL       string
	APIKey        string
	BinanceMarket *BinanceMarket
}

//BinanceMarket has information for market related binance endpoints
type BinanceMarket struct {
	Depth           string
	TradeRecent     string
	TradeHistorical string
}
