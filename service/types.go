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

type Coinbase struct {
	BaseURL                string
	APIKey                 string
	SecretKey              string
	CoinbaseUser           *CoinbaseUser
	CoinbaseAccount        *CoinbaseAccount
	CoinbaseAddresses      *CoinbaseAddresses
	CoinbaseTransactions   *CoinbaseTransactions
	CoinbaseBuys           *CoinbaseBuys
	CoinbaseSells          *CoinbaseSells
	CoinbaseDeposits       *CoinbaseDeposits
	CoinbaseWithdrawals    *CoinbaseWithdrawals
	CoinbasePaymentMethods *CoinbasePaymentMethods
}

type CoinbaseUser struct {
	UserEndpoint string
}

type CoinbaseAccount struct {
	AccountEndpoint string
}

type CoinbaseAddresses struct {
	AddressEndpoint string
}

type CoinbaseTransactions struct {
	TransactionsEndpoints string
}

type CoinbaseBuys struct {
	BuysEndpoint string
}

type CoinbaseSells struct {
	SellEndpoint string
}

type CoinbaseDeposits struct {
	DepositEndpoint string
}

type CoinbaseWithdrawals struct {
	WithdrawalsEndpoint string
}

type CoinbasePaymentMethods struct {
	PaymentMethodEndpoint string
}
