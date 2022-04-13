package core

import (
	"encoding/json"
	"strconv"
)

func MapToJson(toMap interface{}) string {
	result, err := json.Marshal(toMap)
	if err != nil {
		return "{}"
	}

	return string(result)
}

func MapFromEscapedJson(toMap string, target interface{}) error {
	result, err := strconv.Unquote(`"` + toMap + `"`)
	if err != nil {
		return err
	}

	return MapFromJson(result, target)
}

func MapFromJson(toMap string, target interface{}) error {
	err := json.Unmarshal([]byte(toMap), target)

	return err
}
