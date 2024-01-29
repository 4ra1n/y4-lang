package test

import (
	"bytes"
	"testing"

	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
)

func TestCase28(t *testing.T) {
	code := `
// function binary_search_iterative
// this function performs a binary search on a sorted array using iteration
// parameters:
// arr: The sorted array in which the search is performed
// target: The value to search for
// returns: The index of the target element if found, or -1 if not found
def binary_search_iterative(arr, target) {
    low = 0;
    high = length(arr) - 1;
    while low <= high {
        // calculate the middle index
        mid = low + (high - low) / 2;
        // check if the middle element is the target
        if arr[mid] == target {
            return mid;
        }
        // if target is smaller than mid, then it can only be present in the left subarray
        if arr[mid] > target {
            high = mid - 1;
        } else {
            // otherwise the element can only be present in the right subarray
            low = mid + 1;
        }
    }
    // target is not present in the array

}

arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21];
target = 9;
index_recursive := binary_search_iterative(arr, target);
print(index_recursive);


`
	log.SetLevel(log.DebugLevel)
	r := bytes.NewReader([]byte(code))
	i := core.NewInterpreter(lexer.NewLexer(r), nil)
	i.Start()
}