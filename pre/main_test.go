package pre

import (
	"testing"

	"github.com/4ra1n/y4-lang/assert"
)

func TestSearchMain(t *testing.T) {
	file, err := SearchMain([]string{"others.y4", "main.y4"})
	if err != nil {
		return
	}
	assert.Equal(t, file, "main.y4")
}
