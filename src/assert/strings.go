package assert

import "strings"

// ContainsArr checks if string contains multiple sub strings
func ContainsArr(str string, exp []string) bool {
	for _, x := range exp {
		if ContainsString(str, x) == false {
			return false
		}
	}
	return true
}

// EqualsArr compares a string to every single element of a list
func EqualsArr(str string, exp []string) bool {
	for _, x := range exp {
		if (str == x) == false {
			return false
		}
	}
	return true
}

// ContainsString checks if string contains a substring
func ContainsString(str, exp string) bool {
	return strings.Contains(str, exp)
}

// EqualsString compares two strings
func EqualsString(str, exp string) bool {
	return str == exp
}
