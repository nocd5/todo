package common

import (
	"time"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetModifiedTime() string {
	return time.Now().UTC().Truncate(time.Millisecond).Format(time.RFC3339Nano)
}
