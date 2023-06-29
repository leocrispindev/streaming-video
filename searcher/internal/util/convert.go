package util

import (
	"encoding/json"
	"strconv"
)

func ConvertToInt(value interface{}) int {
	str, _ := json.Marshal(value)

	valueInt, _ := strconv.Atoi(string(str))

	return valueInt
}
