package lib

import (
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
