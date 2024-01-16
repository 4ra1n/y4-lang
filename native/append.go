package native

const nativeAppendFunction = "append"

func y4Append(arr []interface{}, val interface{}) []interface{} {
	return append(arr, val)
}
