package general

import (
	"encoding/json"
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/config"
	"github.com/ivandonyk/Crypto-Trader/lib"
	"net/http"
)

const (
	EndpointPing         = "/api/v3/ping"
	EndpointTime         = "/api/v3/time"
	EndpointExchangeInfo = "/api/v3/exchangeInfo"
)

type GenResp func() (*PingResponse, error)

type General struct {
	baseURL         string
	pingRequestFunc func() (*PingResponse, error)
}

type checkServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type PingResponse struct {
}

func NewGeneralAPI(c *config.Config) (*General, error) {

	return &General{
		baseURL: c.BaseURL,
	}, nil

}

func (g *General) GetPing() (*PingResponse, error) {

	var pingResp *PingResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, EndpointPing), nil)
	if err != nil {
		return nil, fmt.Errorf("error building new request object")
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occured when attempting to build request %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&pingResp); err != nil {
		return nil, fmt.Errorf("error decoding new json value %v", err)
	}

	return pingResp, nil
}

func (g *General) CheckServiceTime() (*checkServerTimeResponse, error) {
	var checkServerResponse *checkServerTimeResponse

	req, err := http.NewRequest(http.MethodGet, lib.URLJoin(g.baseURL, EndpointTime), nil)
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

func (p *PingResponse) ToJson() (string, error) {

	return lib.ToJson(p)

}

func (cs *checkServerTimeResponse) ToJson() (string, error)  {
	return lib.ToJson(cs)

}
