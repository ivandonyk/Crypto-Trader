package main

import (
	"github.com/ivandonyk/Crypto-Trader/client/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	var Version string
	app := cli.NewApp()
	app.Version = Version
	app.Author = "Brandon Harrell"
	app.Name = "crypt-trader"
	app.Usage = "crypto-trader <exchange> action"
	app.Commands = []cli.Command{
		cmd.BinanceCmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
