package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"webra/src/assert"
	"webra/src/request"
)

type tChanProcess chan tTestCase
type tChanDone chan bool

func (wra *tWebRA) processTestSuite() {
	chProcess := make(tChanProcess, CLI.Threads)
	chDone := make(tChanDone, CLI.Threads)

	for _, testcase := range wra.TestSuite {
		go wra.processTestCase(testcase, chProcess, chDone)
	}

	counter := 0
	var newTestSuite []tTestCase
	for _ = range chDone {
		counter++
		newTestSuite = append(newTestSuite, <-chProcess)
		if counter >= len(wra.TestSuite) {
			close(chProcess)
			close(chDone)
			break
		}
	}
	sort.Sort(tTestSuite(newTestSuite))
	wra.TestSuite = newTestSuite
}

func (wra *tWebRA) processTestCase(testcase tTestCase, chProcess tChanProcess, chDone tChanDone) {
	req := request.Init(CLI.UserAgent, CLI.Timeout)
	resp, responseErr := req.HTTP(testcase.URL)

	testcase.Tests[0].Result.Success = true
	if responseErr != nil {
		testcase.Tests[0].Result.Success = false
		testcase.Tests[0].Result.Msg = wra.addTestMessage(
			testcase.Tests[0], "could not connect: %q", responseErr,
		)
	} else {
		for idx, test := range testcase.Tests {
			if test.Name == "StatusCodeEquals" {
				test.Result = wra.assertEquals(strconv.Itoa(resp.StatusCode), test)
				test.Result.Msg = wra.addTestMessage(
					test, "status code was %d not %s",
					resp.StatusCode, test.Expectations)
			}
			testcase.Tests[idx] = test

			if test.Name == "BodyContains" {
				test.Result = wra.assertContains(ioToString(resp.Body), test)

				test.Result.Msg = wra.addTestMessage(
					test, "body did not contain %s", test.Expectations,
				)
			}
			testcase.Tests[idx] = test

			if test.Name == "HeaderKeyVal" {
				test.Result = wra.assertHeader(resp.Header, test)
				test.Result.Msg = wra.addTestMessage(
					test, "header key val not %s", test.Expectations,
				)
			}
			testcase.Tests[idx] = test
		}
	}
	testcase.Result.Msg = wra.makeTestCaseErrMessage(testcase)
	testcase.Result.Success = true
	if len(testcase.Result.Msg) > 0 {
		testcase.Result.Success = false
	}
	chProcess <- testcase
	chDone <- true
}

func (wra *tWebRA) assertContains(s string, test tTest) (res tResult) {
	res.Success = assert.ContainsArr(s, test.Expectations)
	return
}

func (wra *tWebRA) assertEquals(s string, test tTest) (res tResult) {
	res.Success = assert.EqualsArr(s, test.Expectations)
	return
}

func (wra *tWebRA) assertHeader(header http.Header, test tTest) (res tResult) {
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

func (wra *tWebRA) makeTestCaseErrMessage(testcase tTestCase) (s string) {
	var arr []string
	for _, test := range testcase.Tests {
		if test.Result.Msg != "" {
			arr = append(arr, test.Result.Msg)
		}
		s = strings.Join(arr, "; ")
	}
	return
}

func (wra *tWebRA) addTestMessage(test tTest, msg string, itf ...interface{}) (s string) {
	if test.Result.Success == false {
		if len(itf) == 1 {
			s = fmt.Sprintf(msg, itf[0])
		}
		if len(itf) == 2 {
			s = fmt.Sprintf(msg, itf[0], itf[1])
		}
	}
	return
}
