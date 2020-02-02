package main

import (
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/config"
	"github.com/ivandonyk/Crypto-Trader/general"
	"log"
)

func main() {

	c := config.Config{
		BaseURL: "https://api.binance.com",
	}

	ping, err := general.NewGeneralAPI(&c)
	if err != nil{
		log.Fatal(err)
	}

	resp, err := ping.GetPing()
	if err != nil{
		log.Fatal(err)
	}

	pingJson, err := resp.ToJson()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(pingJson)

	serverTime, err := ping.CheckServiceTime()
	if err != nil{
		log.Fatal(err)
	}

	serverTimeJson, err := serverTime.ToJson()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(serverTimeJson)


}
