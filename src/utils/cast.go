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

func StringToBool(s string) (bool, error) {
	if s == "false" {
		return false, nil
	} else if s == "true" {
		return true, nil
	}
	return false, errors.New("params error: should be true or false")
}
