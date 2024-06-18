package native

import "fmt"

const nativePrintFunction = "打印"

func y4Print(v interface{}) {
	fmt.Println(v)
}
