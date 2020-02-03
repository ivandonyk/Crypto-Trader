package cmd

import (
	"fmt"
	"github.com/brharrelldev/crytoTrader/config"
	"github.com/brharrelldev/crytoTrader/general"
	"github.com/urfave/cli"
)

const (
	baseURL = "https://api.binance.com"
)

//Binance commands
var BinanceCmd cli.Command

func init() {

	BinanceCmd.Name = "binance"
	BinanceCmd.Usage = "binance <action>"
	BinanceCmd.Subcommands = []cli.Command{
		{
			Name:    "server-time",
			Aliases: []string{"st"},
			Action:  serverTimeAction,
		},
		{
			Name:   "ping",
			Action: pingAction,
		},
	}
}

func pingAction(c *cli.Context) error {
	conf := &config.Config{
		BaseURL: baseURL,
	}
	ping, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	pingResp, err := ping.GetPing()
	if err != nil {
		return fmt.Errorf("error occurred when checking server time %v", err)
	}

	pingJson, err := pingResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(pingJson)

	return nil
}

func serverTimeAction(c *cli.Context) error {
	conf := &config.Config{
		BaseURL: baseURL,
	}

	st, err := general.NewGeneralAPI(conf)
	if err != nil {
		return fmt.Errorf("could not instantiate general api")
	}

	stResp, err := st.CheckServiceTime()
	if err != nil {
		return fmt.Errorf("error occurred when checking server time %v", err)
	}

	stJson, err := stResp.ToJson()
	if err != nil {
		return fmt.Errorf("error serializing json object %v", err)
	}

	fmt.Println(stJson)

	return nil

}
