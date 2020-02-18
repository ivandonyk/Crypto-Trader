package service

import (
	"context"
	api_coinbase "github.com/ivandonyk/Crypto-Trader/api/coinbase"
	"github.com/ivandonyk/Crypto-Trader/exchanges/coinbase/Users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ShowCurrentUser(ctx context.Context, request *api_coinbase.ShowUser) (*api_coinbase.ShowUserResponse, error) {
	currentUser, err := Users.ShowCurrentUser(s.Coinbase.BaseURL, s.Coinbase.CoinbaseUser.UserEndpoint,
		s.Coinbase.APIKey, s.Coinbase.SecretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get current user %v", err)
	}

	return &api_coinbase.ShowUserResponse{
		Data: &api_coinbase.User{
			Id:           currentUser.Data.ID,
			Name:         currentUser.Data.Name,
			UserName:     currentUser.Data.UserName,
			ProfileUrl:   currentUser.Data.ProfileURL,
			AvatarUrl:    currentUser.Data.AvatarURL,
			Resource:     currentUser.Data.Resource,
			ResourcePath: currentUser.Data.ResourcePath,
		},
	}, nil
}
