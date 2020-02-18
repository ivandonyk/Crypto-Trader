package cmd

import (
	"context"
	"fmt"
	api_coinbase "github.com/ivandonyk/Crypto-Trader/api/coinbase"
	"github.com/ivandonyk/Crypto-Trader/build_tags"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"
)

var CoinbaseCmd cli.Command

func init() {
	CoinbaseCmd.Name = "coinbase"
	CoinbaseCmd.Aliases = []string{"cb"}
	CoinbaseCmd.Subcommands = []cli.Command{
		{
			Name:    "current_user",
			Aliases: []string{"cu"},
			Action:  ShowCurrentUser,
		},
	}
}

func ShowCurrentUser(cli *cli.Context) error {

	conn, err := grpc.Dial(build_tags.GRPCHost+":"+build_tags.Port, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("error dialing grpc server %v", err)
	}

	client := api_coinbase.NewCoinbaseUsersClient(conn)

	resp, err := client.ShowCurrentUser(context.Background(), &api_coinbase.ShowUser{})
	if err != nil {
		return fmt.Errorf("error getting response %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "user_name", "avatar_url", "resource", "resource_path"})
	table.Append([]string{resp.Data.Id, resp.Data.Name, resp.Data.UserName, resp.Data.AvatarUrl, resp.Data.Resource,
		resp.Data.ResourcePath})
	table.SetColWidth(100)

	table.Render()

	defer conn.Close()

	return nil

}
