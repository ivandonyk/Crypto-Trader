package service

import (
	"fmt"
	api_binance "github.com/ivandonyk/Crypto-Trader/api/binance"
	"github.com/ivandonyk/Crypto-Trader/ct_config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Binance *Binance
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
	}
}

//StartgRPC this start the grpc server, and registers RPCs
func (s *Server) StartgRPC() error {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("error building new listner %v", err)
	}

	server := grpc.NewServer()

	api_binance.RegisterBinanceMarketDataServer(server, s)

	return server.Serve(lis)
}
