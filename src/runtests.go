package main

import (
	"fmt"
	"webra/src/assert"
)

func (wra *tWebRA) runTests() {
	for _, el := range wra.TestsIterator {
		test := wra.TestsMap[el]
		resp := request(test.URL)

		if test.ExpStatusCode != 0 {
			succ := test.ExpStatusCode == resp.StatusCode
			res := tResult{
				Success: succ, Got: resp.StatusCode, Exp: test.ExpStatusCode,
			}
			res.Msg = fmt.Sprintf("good, got %d", resp.StatusCode)
			if succ == false {
				res.Msg = fmt.Sprintf(
					"status code, %d instead of %d",
					resp.StatusCode, test.ExpStatusCode,
				)
			}
			test.Result.List = append(test.Result.List, res)
		}

		if test.ExpBodySubstrings != nil {
			succ := assert.Substrings(
				ioToString(resp.Body), test.ExpBodySubstrings,
			)
			res := tResult{
				Success: succ,
				Got:     resp.Body,
				Exp:     test.ExpStatusCode,
			}
			res.Msg = fmt.Sprintf(
				"good, got %q", test.ExpBodySubstrings,
			)
			if succ == false {
				res.Msg = fmt.Sprintf(
					"body did not contain %q", test.ExpBodySubstrings,
				)
			}
			test.Result.List = append(test.Result.List, res)
		}

		test.Result.Success = testSuccessful(test)
		test.Result.Msg = finalTestMessage(test)
		wra.TestsMap[el] = test
	}
}

func testSuccessful(test tTest) bool {
	for _, res := range test.Result.List {
		if res.Success == false {
			return false
		}
	}
	return true
}

func finalTestMessage(test tTest) string {
	var arr []string
	for _, res := range test.Result.List {
		arr = append(arr, res.Msg)
	}
	return fmt.Sprintf("%s", arr)
}
