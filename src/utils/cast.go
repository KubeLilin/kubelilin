package utils

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"strconv"
	"time"
)

// Int32Ptr convert int32 value to a pointer.
func Int32Ptr(i int32) *int32 {
	return &i
}

func PtrToInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func PtrToInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// Int64Ptr convert int64 value to a pointer.
func Int64Ptr(i int64) *int64 {
	return &i
}

// StringPtr convert string value to a pointer.
func StringPtr(s string) *string {
	return &s
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToUInt64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToBool(s string) (bool, error) {
	if s == "false" {
		return false, nil
	} else if s == "true" {
		return true, nil
	}
	return false, errors.New("params error: should be true or false")
}

func StringToNull(s string) interface{} {
	if s == "" {
		return nil
	} else {
		return s
	}
}

func Uint8ToBool(num uint8) bool {
	if num > 0 {
		return true
	} else {
		return false
	}
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case *uint:
		return strconv.FormatUint(uint64(*v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}

func ObjectToString(obj interface{}) string {
	jsonByte, _ := json.Marshal(obj)
	return string(jsonByte)
}

func StringToJson(jsonStr string, obj any) {
	_ = json.Unmarshal([]byte(jsonStr), obj)
}

// yaml to json
func YamlToJson(yamlString string) (string, error) {
	var data map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}
	jsonString, _ := json.Marshal(data)
	return string(jsonString), nil
}

// yaml to object
func YamlToObject(yamlString string, obj any) error {
	err := yaml.Unmarshal([]byte(yamlString), obj)
	if err != nil {
		return err
	}
	return nil
}

func TimeFormat(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
