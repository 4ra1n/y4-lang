package native

import (
	"time"
)

const nativeNowFunction = "now"

func y4Now() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
