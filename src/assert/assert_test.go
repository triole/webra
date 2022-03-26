package assert

import (
	"fmt"
	"testing"
)

func TestSubstrings(t *testing.T) {
	runTestSubstrings("hello world", []string{"hello"}, true, t)
	runTestSubstrings("hello world", []string{"hello", "world"}, true, t)
	runTestSubstrings("hello world", []string{"hello", "space"}, false, t)
}

func runTestSubstrings(str string, substr []string, exp bool, t *testing.T) {
	res := Substrings(str, substr)
	if res != exp {
		fmt.Printf(
			"Assert substring failed: %s, %q was %v not %v\n",
			str, substr, res, exp,
		)
	}
}
