package native

const nativeIntFunction = "转整数"

func y4Int(i interface{}) int {
	v, ok := i.(int)
	if ok {
		return v
	} else {
		return 0
	}
}
