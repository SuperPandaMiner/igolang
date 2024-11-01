package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

func ParseUint64(s string, bitSize int) (uint64, error) {
	return strconv.ParseUint(s, 10, bitSize)
}

func ParseInt64(s string, bitSize int) (int64, error) {
	return strconv.ParseInt(s, 10, bitSize)
}

func ToJsonString(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	} else {
		jsonString := string(jsonData)
		return jsonString, nil
	}
}

func JsonToObject[T any](jsonString string) (*T, error) {
	var result T
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func JsonToMap(jsonString string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ParseToStandardTime(t time.Time) string {
	return t.Format(time.DateTime)
}
