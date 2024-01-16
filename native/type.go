package native

const nativeTypeFunction = "type"

func y4Type(v interface{}) string {
	_, isS := v.(string)
	if isS {
		return "<string>"
	}
	_, isI := v.(int)
	if isI {
		return "<int>"
	}
	_, isF := v.(float64)
	if isF {
		return "<float>"
	}
	return "<object>"
}
