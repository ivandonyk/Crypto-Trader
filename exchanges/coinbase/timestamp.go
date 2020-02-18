package coinbase

import (
	"encoding/json"
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/lib"
	"net/http"
)

type Data struct {
	Iso   string `json:"iso"`
	Epoch int64  `json:"epoch"`
}
type TimeStamp struct {
	Data Data `json:"data"`
}

func GetTimestamp(baseUrl string) (int64, error) {
	var ts TimeStamp

	reqURL := lib.URLJoin(baseUrl, "/v2/time")
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, fmt.Errorf("error getting timestamp %v", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error getting response %v", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&ts); err != nil {
		return 0, fmt.Errorf("error serializing json response %v", err)
	}

	return ts.Data.Epoch, nil
}
