package assert

import "strings"

func Substrings(str string, exp []string) bool {
	for _, x := range exp {
		if strings.Contains(str, x) == false {
			return false
		}
	}
	return true
}
