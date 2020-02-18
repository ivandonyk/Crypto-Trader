package ct_config

//Config is a generic config struct.  More will be added later
type Config struct {
	BinanceConfig  *BinanceConfig
	CoinbaseConfig *CoinbaseConfig
}

//BinanceConfig is the main config object for the Binance exchange
type BinanceConfig struct {
	BaseURL          string
	APIKey           string
	Depth            string
	TradeRecent      string
	TradeHistorical  string
	AggregateTrades  string
	CandleStick      string
	CurrentAvgPrice  string
	DayStat          string
	PriceTicker      string
	BookTicker       string
	NewOrder         string
	NewOrderTest     string
	QueryOrder       string
	CancelOrder      string
	CurrentOpenOrder string
	AllOrders        string
	NewOCO           string
	CancelOCO        string
	QueryOCO         string
	AccountInfo      string
	AccountTradeList string
	StartUserStream  string
}

type CoinbaseConfig struct {
	BaseURL        string
	APIKey         string
	APISecret      string
	User           string
	Accounts       string
	Addresses      string
	Transactions   string
	Buys           string
	Sells          string
	Deposits       string
	Withdrawals    string
	PaymentMethods string
}
