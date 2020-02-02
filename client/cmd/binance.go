package cmd

import (
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/config"
	"github.com/ivandonyk/Crypto-Trader/general"
	"github.com/urfave/cli"
)

const (
	BaseURL = "https://api.binance.com"
)

var BinanceCmd cli.Command
var conf *config.Config

func init() {
	conf = &config.Config{
		BaseURL: BaseURL,
	}

	BinanceCmd.Name = "binance"
	BinanceCmd.Usage = "binance <action>"
	BinanceCmd.Subcommands = []cli.Command{
		{
			Name:    "server-time",
			Aliases: []string{"st"},
			Action:  ServerTimeAction,
		},
		{
			Name:   "ping",
			Action: PingAction,
		},
	}
}

func PingAction(c *cli.Context) error {
	ping, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	pingResp, err := ping.GetPing()
	if err != nil {
		return fmt.Errorf("error occured when checking server time %v", err)
	}

	pingJson, err := pingResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(pingJson)

	return nil
}

func ServerTimeAction(c *cli.Context) error {

	st, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	stResp, err := st.CheckServiceTime()
	if err != nil {
		return fmt.Errorf("error occured when checking server time %v", err)
	}

	stJson, err := stResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(stJson)

	return nil

}
