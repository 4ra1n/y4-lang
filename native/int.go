package native

const nativeIntFunction = "int"

func y4Int(i interface{}) int {
	v, ok := i.(int)
	if ok {
		return v
	} else {
		return 0
	}
}
