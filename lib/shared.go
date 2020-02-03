package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)

//common method for joining url.  I'm sure there is a library for this
func URLJoin(base, endpoint string) string {

	return strings.Join([]string{base, endpoint}, "")
}

//common json encoding method to reduce repetitive code
func ToJson(resp interface{}) (string, error) {

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return "", fmt.Errorf("could not create new json object %v", err)
	}

	return string(jsonBytes), nil
}
