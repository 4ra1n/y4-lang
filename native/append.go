package native

const nativeAppendFunction = "追加"

func y4Append(arr []interface{}, val interface{}) []interface{} {
	return append(arr, val)
}
