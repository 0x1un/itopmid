package util

import "time"

func FormatTimeToString(time time.Time) string {
	return time.Local().Format("15:04:05")
}
