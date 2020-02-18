package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

//URLJoin is a helper method to help concatenate baseURL and endpoint
func URLJoin(base, endpoint string) string {

	return strings.Join([]string{base, endpoint}, "")
}

//ToJson common ToJson method to be included with other formatters.
func ToJson(resp interface{}) (string, error) {

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("could not create new json object %v", err)
	}

	return string(jsonBytes), nil
}

func CoinbaseAuthorize(epoch, secretKey, method string, requestPath string) string {
	fmt.Println(requestPath)
	message := strings.Join([]string{epoch, method, requestPath}, "")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))

	signature := hex.EncodeToString(mac.Sum(nil))

	return signature

}
