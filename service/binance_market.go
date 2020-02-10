package service

import (
	"context"
	api_binance "github.com/ivandonyk/Crypto-Trader/api/binance"
	"github.com/ivandonyk/Crypto-Trader/exchanges/binance_api/market"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//getBinanceMarket depth.  This takes in symbol and limit to get the market depth.  Max limit is 5000
func (s *Server) GetBinanceMarketDepth(ctx context.Context, req *api_binance.GetBinanceMarketDepthRequest) (*api_binance.GetBinanceMarketDepthResponse, error) {
	symbol := req.Symbol
	limit := req.Limit

	baseURL := s.Binance.BaseURL
	endpoint := s.Binance.BinanceMarket.Depth

	p := market.ParamsMarket{
		Symbol: symbol,
		Limit:  limit,
	}

	marketData, err := market.GetMarketDepth(baseURL, endpoint, p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting market depth %v", err)
	}

	var asksSlice []*api_binance.Asks
	var bidsSlice []*api_binance.Bids
	//
	for _, asks := range marketData.Ask {
		askPrice := &api_binance.Asks{
			Low:  asks[0],
			High: asks[1],
		}
		asksSlice = append(asksSlice, askPrice)

	}

	for _, bids := range marketData.Bid {
		bidsSlice = append(bidsSlice, &api_binance.Bids{
			Low:  bids[0],
			High: bids[1],
		})
	}

	return &api_binance.GetBinanceMarketDepthResponse{
		LastUpdateId: marketData.ID,
		Bids:         bidsSlice,
		Asks:         asksSlice,
	}, nil
}

//GetBinanceMarketTradesRecent.  This takes symbol.  This gets the most recent trades
func (s *Server) GetBinanceMarketTradesRecent(ctx context.Context, req *api_binance.GetBinanceMarketTradesRecentRequest) (*api_binance.GetBinanceMarketTradesRecentResponse, error) {
	symbol := req.Symbol
	limit := req.Limit

	baseURL := s.Binance.BaseURL
	endpoint := s.Binance.BinanceMarket.TradeRecent

	p := market.ParamsMarket{
		Symbol: symbol,
		Limit:  limit,
	}

	recentTrades, err := market.GetRecentTrades(baseURL, endpoint, p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting recent trade data %v", err)
	}

	var tiSlice []*api_binance.TradeInfo
	for rt := range recentTrades.Iterate() {
		ti := &api_binance.TradeInfo{
			ID:            rt.ID,
			Time:          rt.Time,
			Price:         rt.Price,
			Quantity:      rt.Qty,
			QuoteQuantity: rt.QuoteQty,
			IsBuyerMake:   rt.IsBuyerMarker,
			IsBestMatch:   rt.IsBestMatch,
		}

		tiSlice = append(tiSlice, ti)
	}

	return &api_binance.GetBinanceMarketTradesRecentResponse{
		Results: tiSlice,
	}, nil

}

//GetBinanceMarketTradesHistorical.  These are historical trades.  You can also search these trades by tradeID
func (s *Server) GetBinanceMarketTradesHistorical(ctx context.Context, req *api_binance.GetBinanceMarketTradesHistoricalRequest) (*api_binance.GetBinanceMarketTradesHistoricalResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
