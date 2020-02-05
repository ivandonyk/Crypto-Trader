package ct_config

//Config is a generic config struct.  More will be added later
type Config struct {
	BinanceConfig *BinanceConfig
}

//BinanceConfig is the main config object for the Binance exchange
type BinanceConfig struct {
	BaseURL         string
	APIKey          string
	Depth           string
	TradeRecent     string
	TradeHistorical string
}
