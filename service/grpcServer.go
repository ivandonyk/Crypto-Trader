package service

import (
	"fmt"
	api_binance "github.com/ivandonyk/Crypto-Trader/api/binance"
	api_coinbase "github.com/ivandonyk/Crypto-Trader/api/coinbase"
	"github.com/ivandonyk/Crypto-Trader/ct_config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Binance  *Binance
	Coinbase *Coinbase
}

//NewServer starts a new server object.  This is the main object to coordinate the gRPC calls
func NewServer(c ct_config.Config) *Server {

	return &Server{
		Binance: &Binance{
			BaseURL: c.BinanceConfig.BaseURL,
			BinanceMarket: &BinanceMarket{
				Depth:           c.BinanceConfig.Depth,
				TradeRecent:     c.BinanceConfig.TradeRecent,
				TradeHistorical: c.BinanceConfig.TradeHistorical,
			},
		},
		Coinbase: &Coinbase{
			BaseURL:   c.CoinbaseConfig.BaseURL,
			APIKey:    c.CoinbaseConfig.APIKey,
			SecretKey: c.CoinbaseConfig.APISecret,
			CoinbaseUser: &CoinbaseUser{
				UserEndpoint: c.CoinbaseConfig.User,
			},
		},
	}
}

//StartgRPC this start the grpc server, and registers RPCs
func (s *Server) StartgRPC() error {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return fmt.Errorf("error building new listner %v", err)
	}

	server := grpc.NewServer()

	api_binance.RegisterBinanceMarketDataServer(server, s)
	api_coinbase.RegisterCoinbaseUsersServer(server, s)

	return server.Serve(lis)
}
