package utils

import "encoding/json"

type JsonUtil struct{}

func (ju *JsonUtil) Stringify(input interface{}) string {
	jsonified, _ := json.Marshal(input)
	return string(jsonified)
}
