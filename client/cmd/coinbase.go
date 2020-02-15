package cmd

import "github.com/urfave/cli"

var CoinbaseCmd cli.Command

func init() {
	CoinbaseCmd.Name = "coinbase"
	CoinbaseCmd.Aliases = []string{"cb"}
	CoinbaseCmd.Subcommands = []cli.Command{
		{
			Name: "current_user",
			Aliases: []string{"cu"},
			Action: ShowCurrentUser,
		},
	}
}

func ShowCurrentUser(cli *cli.Context) error  {

	return nil

}
