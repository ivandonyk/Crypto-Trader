package main

import (
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/ct_config"
	"github.com/ivandonyk/Crypto-Trader/service"
	"log"
	"os"
	"os/signal"
)

func main() {
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

	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		fmt.Println("stopping server")
		return
	case err := <-errChan:
		log.Fatalf("error occured during run %v", err)

	}

}
