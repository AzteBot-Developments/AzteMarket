package utils

import "time"

func FormatUnixAsString(timestamp int64, format string) string {

	var ts time.Time
	var timeString string

	ts = time.Unix(timestamp, 0).UTC()
	timeString = ts.Format(format) // e.g -> "Mon, 02 Jan 2006 15:04:05 MST"

	return timeString

}
