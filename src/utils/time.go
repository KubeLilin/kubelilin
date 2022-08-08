package utils

import "time"

func TimeNowPtr() *time.Time {
	now := time.Now()
	return &now
}

func TimeNow() time.Time {
	return time.Now()
}
