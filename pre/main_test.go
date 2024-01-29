package pre

import (
	"fmt"
	"testing"
)

func TestSearchMain(t *testing.T) {
	file, err := SearchMain([]string{"others.y4", "main.y4"})
	if err != nil {
		return
	}
	fmt.Println(file)
}
