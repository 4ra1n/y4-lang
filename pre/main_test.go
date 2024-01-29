package pre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMain(t *testing.T) {
	file, err := SearchMain([]string{"others.y4", "main.y4"})
	if err != nil {
		return
	}
	assert.Equal(t, file, "main.y4")
}
