package general

import (
	"encoding/json"
	"fmt"
	"github.com/brharrelldev/crytoTrader/config"
	"github.com/brharrelldev/crytoTrader/lib"
	"net/http"
)

const (
	endpointPing         = "/api/v3/ping"
	endpointTime         = "/api/v3/time"
	endpointExchangeInfo = "/api/v3/exchangeInfo"
)

//General is a struct for finance general API
type General struct {
	baseURL         string
	pingRequestFunc func() (*pingResponse, error)
}

type checkServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type pingResponse struct {
}

//NewGeneralAPI constructor for Binance General API
func NewGeneralAPI(c *config.Config) (*General, error) {

	return &General{
		baseURL: c.BaseURL,
	}, nil

}

//GetPing mirrors the Binance API.  This checks for liveness of the server.  It usually returns an empty response
func (g *General) GetPing() (*pingResponse, error) { //test comment

	var pingResp *pingResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, endpointPing), nil)
	if err != nil {
		return nil, fmt.Errorf("error building new request object")
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred when attempting to build request %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&pingResp); err != nil {
		return nil, fmt.Errorf("error decoding new json value %v", err)
	}

	return pingResp, nil
}

//CheckServiceTime will check the server's time.  Will return response in unix format.  This is set as a int64
func (g *General) CheckServiceTime() (*checkServerTimeResponse, error) {
	var checkServerResponse *checkServerTimeResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, endpointTime), nil)
	if err != nil {
		return nil, fmt.Errorf("error building request %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&checkServerResponse); err != nil {
		return nil, fmt.Errorf("error decoding response %v", err)
	}

	return checkServerResponse, nil

}

//ToJson is to format a json string for pingResponse.
func (p *pingResponse) ToJson() (string, error) {
	return lib.ToJson(p)

}

//ToJson to format to json string for checkServerTimeResponse
func (cs *checkServerTimeResponse) ToJson() (string, error) {
	return lib.ToJson(cs)

}
