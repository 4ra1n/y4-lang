package native

import "time"

const (
	nativeSleepFunction  = "睡眠"
	nativeSleepMFunction = "睡眠毫秒"
)

func y4Sleep(second int) {
	time.Sleep(time.Second * time.Duration(second))
}

func y4SleepM(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}
