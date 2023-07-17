package codetype

import "time"

func StringToTime(format string, input string) (time.Time, error) {
	return time.Parse(format, input)
}
