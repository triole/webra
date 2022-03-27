package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webra/src/assert"
	"webra/src/request"
)

func (wra *tWebRA) runTestSuite() {
	for idx, testcase := range wra.TestSuite {
		wra.TestSuite[idx] = wra.runTestCase(testcase)
	}
}

func (wra *tWebRA) addMessage(test tTest, msg, resp, exp interface{}) (r string) {
	if test.Result.Success == false {
		if resp == nil {
			r = fmt.Sprintf(fmt.Sprintf("%s", msg), exp)
		} else {
			r = fmt.Sprintf(fmt.Sprintf("%s", msg), resp, exp)
		}
	}
	return
}

func (wra *tWebRA) runTestCase(testcase tTestCase) tTestCase {
	req := request.Init(CLI.UserAgent)
	resp, responseErr := req.HTTP(testcase.URL)
	if responseErr != nil {
		testcase.Result.Msg = fmt.Sprintf("%s", responseErr)
		testcase.Result.Success = false
	}
	if responseErr == nil {
		for idx, test := range testcase.Tests {
			if test.Name == "StatusCodeEquals" {
				test.Result = assertEquals(strconv.Itoa(resp.StatusCode), test)
				test.Result.Msg = wra.addMessage(
					test, "status code was %s not %s",
					strconv.Itoa(resp.StatusCode), test.Expectations)
			}
			testcase.Tests[idx] = test

			if test.Name == "BodyContains" {
				test.Result = assertContains(ioToString(resp.Body), test)

				test.Result.Msg = wra.addMessage(
					test, "body did not contain %s", nil, test.Expectations,
				)
			}
			testcase.Tests[idx] = test

			if test.Name == "HeaderKeyVal" {
				test.Result = assertHeader(resp.Header, test)
				test.Result.Msg = wra.addMessage(
					test, "header key val not %s", nil, test.Expectations,
				)
			}
			testcase.Tests[idx] = test

		}
	}
	testcase.Result.Success = false
	testcase.Result.Msg = errMessage(testcase)
	if testcase.Result.Msg == "" {
		testcase.Result.Success = true
	}
	return testcase
}

func assertContains(s string, test tTest) (res tResult) {
	res.Success = assert.ContainsArr(s, test.Expectations)
	return
}

func assertEquals(s string, test tTest) (res tResult) {
	res.Success = assert.EqualsArr(s, test.Expectations)
	return
}

func assertHeader(header http.Header, test tTest) (res tResult) {
	var exp [][]string
	var arr []string
	for _, el := range test.Expectations {
		for _, el := range strings.Split(el, ":") {
			arr = append(
				arr, strings.TrimSpace(el),
			)
		}
		exp = append(exp, arr)
	}
	res.Success = assert.IterHeader(header, exp)
	res.Msg = fmt.Sprintf("assert header failed %s", header)
	return
}

func errMessage(testcase tTestCase) (s string) {
	var arr []string
	for _, test := range testcase.Tests {
		if test.Result.Msg != "" {
			arr = append(arr, test.Result.Msg)
		}
		s = strings.Join(arr, "; ")
	}
	return
}
