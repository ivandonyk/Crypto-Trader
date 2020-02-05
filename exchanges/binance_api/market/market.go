package market

import (
	"encoding/json"
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/lib"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ParamsMarket struct {
	Symbol string
	Limit  int32
	FromID int32
}

type Asks struct {
	Prices []LowHigh
}

type LowHigh struct {
	Low  string
	High string
}

type Bids struct {
	Prices []LowHigh
}

type ResponseMarketDepth struct {
	ID  int32      `json:"lastUpdateId"`
	Ask [][]string `json:"asks"`
	Bid [][]string `json:"bids"`
}

func GetMarketDepth(baseURL string, endpoint string, params ParamsMarket) (*ResponseMarketDepth, error) {
	var marketDepthResponse *ResponseMarketDepth

	reqUrl := lib.URLJoin(baseURL, endpoint)

	rawURL, err := url.Parse(reqUrl)
	if err != nil {
		return marketDepthResponse, nil
	}

	q := rawURL.Query()
	q.Add("symbol", params.Symbol)
	q.Add("limit", strconv.Itoa(int(params.Limit)))

	rawURL.RawQuery = q.Encode()

	fmt.Println(rawURL.String())

	req, err := http.NewRequest(http.MethodGet, rawURL.String(), nil)
	if err != nil {
		return marketDepthResponse, fmt.Errorf("error building new request %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error invoking client %v", err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body %v", err)
	}

	if err := json.Unmarshal(f, &marketDepthResponse); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error serializing json object %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(marketDepthResponse)

	return marketDepthResponse, nil
}
