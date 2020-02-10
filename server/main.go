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
