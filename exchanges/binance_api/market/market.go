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

type TradeInfo struct {
	ID            int32  `json:"id"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	QuoteQty      string `json:"quoteQty"`
	Time          int64  `json:"time"`
	IsBuyerMarker bool   `json:"isBuyerMarker"`
	IsBestMatch   bool   `json:"isBestMatch"`
}
type ResponseRecentTrades []TradeInfo

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

	req, err := http.NewRequest(http.MethodGet, rawURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building new request %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error invoking client %v", err)
	}

	defer resp.Body.Close()

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body %v", err)
	}

	if err := json.Unmarshal(f, &marketDepthResponse); err != nil {
		return nil, fmt.Errorf("error serializing json object %v", err)
	}

	return marketDepthResponse, nil
}

func GetRecentTrades(baseURL, endpoint string, params ParamsMarket) (*ResponseRecentTrades, error) {
	var recentTrades *ResponseRecentTrades

	reqUrl := lib.URLJoin(baseURL, endpoint)

	rawURL, err := url.Parse(reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing url %v", err)
	}

	q := rawURL.Query()
	q.Add("symbol", params.Symbol)
	q.Add("limit", strconv.Itoa(int(params.Limit)))

	rawURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, rawURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error building new request %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error invoking client %v", err)
	}

	defer resp.Body.Close()

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body %v", err)
	}

	if err := json.Unmarshal(f, &recentTrades); err != nil {
		return nil, fmt.Errorf("error serializing json object %v", err)
	}

	return recentTrades, nil
}

func (m ResponseRecentTrades) Iterate() <-chan TradeInfo {
	c := make(chan TradeInfo)
	go func() {
		for _, ti := range m {
			select {
			case c <- ti:
			case <-c:
				close(c)
				return
			}
		}
		close(c)
	}()

	return c
}
