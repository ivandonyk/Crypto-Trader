package main

import (
	"github.com/ivandonyk/Crypto-Trader/ct_config"
	"github.com/ivandonyk/Crypto-Trader/service"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {

	ctLog, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed initialize logs %v", err)
	}

	config := ct_config.Config{
		BinanceConfig: &ct_config.BinanceConfig{
			BaseURL:         "https://api.binance.com",
			Depth:           "/api/v3/depth",
			TradeRecent:     "/api/v3/trades",
			TradeHistorical: "/api/v3/historicalTrades",
		},
		CoinbaseConfig: &ct_config.CoinbaseConfig{
			BaseURL:        "https://api.coinbase.com",
			APIKey:         os.Getenv("CT_COINBASE_API_KEY"),
			APISecret:      os.Getenv("CT_COINBASE_API_SECRET"),
			User:           "/v2/user",
			Accounts:       "/v2/accounts",
			Addresses:      "/v2/accounts/%s/addresses/",
			Transactions:   "/v2/accounts/%s/transactions/",
			Buys:           "/v2/accounts/%s/buys",
			Sells:          "/v2/accounts/%s/sells",
			Deposits:       "/v2/accounts/%s/sells",
			Withdrawals:    "/v2/accounts/%s/withdrawals",
			PaymentMethods: "v2/payment-methods",
		},
	}

	s := service.NewServer(config)
	if s == nil {
		log.Fatal("error instantiating grpc service")
	}

	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)

	go func() {
		if err := s.StartgRPC(); err != nil {
			errChan <- err
		}
		<-sigChan
	}()

	ctLog.Info("server has been started")

	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		ctLog.Warn("server is being stopped")
		return
	case err := <-errChan:
		ctLog.Fatal("error occurred during run of server", zap.Error(err))
	}

}
