package assert

import (
	"fmt"
	"testing"
)

func TestContainsString(t *testing.T) {
	runTestContainsString("hello world", []string{"hello"}, true, t)
	runTestContainsString("hello world", []string{"hello", "world"}, true, t)
	runTestContainsString("hello world", []string{"hello", "space"}, false, t)
}

func runTestContainsString(str string, substr []string, exp bool, t *testing.T) {
	res := ContainsArr(str, substr)
	if res != exp {
		fmt.Printf(
			"Assert substring failed: %s, %q was %v not %v\n",
			str, substr, res, exp,
		)
	}
}
