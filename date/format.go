package date

import "time"

//Format1 for the format of yyyy-MM-dd HH:mm:ss
var Format1 string = "2006-01-02 15:04:05"

//FormatTimeToString format a time with an format
func FormatTimeToString(inputTime time.Time, format string) string {
	return inputTime.Format(format)
}
