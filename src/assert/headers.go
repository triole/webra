package assert

import (
	"net/http"
)

func IterHeader(header http.Header, exp map[string]string) bool {
	for key, val := range exp {
		if headerHasKey(header, key) == false {
			return false
		}
		if val != "" {
			if val != header[key][0] {
				return false
			}
		}
	}
	return true
}

func headerHasKey(header http.Header, key string) bool {
	if _, ok := header[key]; ok {
		return true
	}
	return false
}
