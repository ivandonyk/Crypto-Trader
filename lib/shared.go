package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)

func URLJoin(base, endpoint string) string {

	return strings.Join([]string{base, endpoint}, "")
}

func ToJson(resp interface{}) (string, error) {

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("could not create new json object %v", err)
	}

	return string(jsonBytes), nil

}
