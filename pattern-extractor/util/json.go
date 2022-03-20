package util

import "encoding/json"

func IsJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func Unmarshal(s string) map[string]string {
	var result map[string]string
	json.Unmarshal([]byte(s), &result)

	return result
}
