package assert

import (
	"net/http"
)

func IterHeader(header http.Header, exp [][]string) bool {
	for _, arr := range exp {
		if len(arr) == 1 {
			b := MapHasKey(header, arr[0])
			if b == false {
				return false
			}
		}
		if len(arr) == 2 {
			b := MapHasKeyVal(header, arr[0], arr[1])
			if b == false {
				return false
			}
		}
	}
	return true
}

func MapHasKey(header http.Header, key string) bool {
	return false
}

func MapHasKeyVal(header http.Header, key, val string) bool {
	return header.Get(key) == val
}
