package cmd

import (
	"context"
	"fmt"
	api_binance "github.com/ivandonyk/Crypto-Trader/api/binance"
	"github.com/ivandonyk/Crypto-Trader/ct_config"
	"github.com/ivandonyk/Crypto-Trader/exchanges/binance_api/general"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"os"
)

const (
	baseURL = "https://api.binance.com"
)

//BinanceCmd is for the binance sub-command
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
		{
			Name:    "market-depth",
			Aliases: []string{"md"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "symbol",
					Usage:    "market-depth --symbol=<trading pairs>",
					Required: true,
				},
				cli.IntFlag{
					Name:  "limit",
					Usage: "display depth limit",
					Value: 10,
				},
			},
			Action: getMarketDepth,
		},
	}
}

func pingAction(c *cli.Context) error {
	conf := &ct_config.Config{
		BinanceConfig: &ct_config.BinanceConfig{
			BaseURL: baseURL,
		},
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
	conf := &ct_config.Config{
		BinanceConfig: &ct_config.BinanceConfig{
			BaseURL: baseURL,
		},
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

func getMarketDepth(c *cli.Context) error {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not dial grpc server %v", err)
	}

	defer conn.Close()

	client := api_binance.NewBinanceMarketDataClient(conn)

	getMarketDepth, err := client.GetBinanceMarketDepth(context.Background(), &api_binance.GetBinanceMarketDepthRequest{
		Symbol: c.String("symbol"),
		Limit:  int32(c.Int("limit")),
	})

	if err != nil {
		return fmt.Errorf("error getting market data %v", err)
	}

	var data [][]string
	marketMap := make(map[string]map[string]map[string]string)
	for idx, asks := range getMarketDepth.Asks {
		idxString := string(idx)
		for _, bids := range getMarketDepth.Bids {
			if _, ok := marketMap[idxString]; ok {
				continue
			}
			marketMap[idxString] = map[string]map[string]string{
				"asks": {
					"high": asks.High,
					"low":  asks.Low,
				},
				"bids": {
					"high": bids.High,
					"low":  bids.Low,
				},
			}

		}

		//data = append()
	}

	for _, v := range marketMap {
		askLow := v["asks"]["low"]
		askHigh := v["asks"]["high"]
		bidLow := v["bids"]["low"]
		bidHigh := v["bids"]["high"]
		record := []string{askLow, askHigh, bidLow, bidHigh}
		data = append(data, record)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ask_low", "ask_high", "bid_low", "bid_high"})

	for _, info := range data {
		table.Append(info)
	}

	table.Render()

	return nil

}
