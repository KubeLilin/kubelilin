package utils

import (
	"errors"
	"strconv"
)

// Int32Ptr convert int32 value to a pointer.
func Int32Ptr(i int32) *int32 {
	return &i
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
