package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {

	baseURL := "http://localhost/someendpoint"

	rawUrL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	q := rawUrL.Query()
	q.Add("symbol", "LTCBTC")

	rawUrL.RawQuery = q.Encode()

	fmt.Println(rawUrL)

}
