package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/stretchr/testify/assert"
)

func TestCase28(t *testing.T) {
	Finish()
	Redirect()
	code := `
// function binary_search_iterative
// this function performs a binary search on a sorted array using iteration
// parameters:
// arr: The sorted array in which the search is performed
// target: The value to search for
// returns: The index of the target element if found, or -1 if not found
函数 binary_search_iterative(arr, target) {
    low = 0;
    high = 长度(arr) - 1;
    当 low <= high {
        // calculate the middle index
        mid = low + (high - low) / 2;
        // check if the middle element is the target
        如果 arr[mid] == target {
            返回 mid;
        }
        // if target is smaller than mid, then it can only be present in the left subarray
        如果 arr[mid] > target {
            high = mid - 1;
        } 另外 {
            // otherwise the element can only be present in the right subarray
            low = mid + 1;
        }
    }
    // target is not present in the array

}

arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21];
target = 9;
index_recursive = binary_search_iterative(arr, target);
打印(index_recursive);


`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()

	result := Read()
	assert.Contains(t, result, "4\n")
}
