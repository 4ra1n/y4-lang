package native

import "time"

const nativeSleepFunction = "sleep"

func y4Sleep(second int) {
	time.Sleep(time.Second * time.Duration(second))
}
