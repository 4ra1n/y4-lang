package native

import "fmt"

const nativePrintFunction = "print"

func y4Print(v interface{}) {
	fmt.Println(v)
}
