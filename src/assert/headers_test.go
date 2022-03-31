package assert

import (
	"fmt"
	"net/http"
	"testing"
)

func makeDummyHeader() (hdr http.Header) {
	hdr = make(http.Header)
	hdr["Content-Length"] = []string{"171"}
	hdr["Content-Type"] = []string{"text/plain"}
	hdr["Date"] = []string{"Thu, 31 Mar 2022 21:18:56 GMT"}
	return
}

func TestIterHeader(t *testing.T) {
	AssertIterHeader(
		map[string]string{"Date": "", "Content-Type": ""},
		true, t)
	AssertIterHeader(
		map[string]string{"Content-Type": "text/plain"},
		true, t)
	AssertIterHeader(
		map[string]string{"Content-Type": "notext/bizarre"},
		false, t)
}

func AssertIterHeader(exp map[string]string, testExp bool, t *testing.T) {
	hdr := makeDummyHeader()
	if IterHeader(hdr, exp) != testExp {
		fmt.Printf("Iter header failed: %q does not contain %q\n", hdr, exp)
	}
}
