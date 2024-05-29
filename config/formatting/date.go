package formatting

import (
	"time"
)

func UnixToHumanDate(unixTimestamp int64) string {
	// Konversi UNIX timestamp ke objek waktu
	timeObj := time.Unix(unixTimestamp, 0)

	return timeObj.Format("02 January 2006 15:03:04")
}

func TimeToHumanDate(timeObj time.Time) string {
	return timeObj.Format("02 January 2006 15:03:04")
}

func TimeParseLocalTz(timeString string, format string) (timeObj time.Time, err error) {

	localTz, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return
	}
	timeObj, err = time.ParseInLocation(format, timeString, localTz)
	return
}
