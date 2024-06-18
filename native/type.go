package native

const nativeTypeFunction = "类型"

func y4Type(v interface{}) string {
	_, isS := v.(string)
	if isS {
		return "<字符串>"
	}
	_, isI := v.(int)
	if isI {
		return "<整数>"
	}
	_, isF := v.(float64)
	if isF {
		return "<浮点数>"
	}
	return "<对象>"
}
