package utils

import "encoding/json"

func GetJsonString(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}
