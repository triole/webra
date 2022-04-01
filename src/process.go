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
	req := request.Init(wra.Settings)
	resp, responseErr := req.HTTP(testcase.URL)

	testcase.Assertions[0].Result.Success = true
	if responseErr != nil {
		testcase.Assertions[0].Result.Success = false
		testcase.Assertions[0].Result.Msg = wra.addTestMessage(
			testcase.Assertions[0], "could not connect: %q", responseErr,
		)
	} else {
		for idx, ase := range testcase.Assertions {
			if ase.Name == "status_code_equals" {
				ase.Result.Success = assert.EqualsArr(
					strconv.Itoa(resp.StatusCode), ase.Expectations,
				)
				ase.Result.Msg = wra.addTestMessage(
					ase, "status code was %d not %s",
					resp.StatusCode, ase.Expectations)
			}
			testcase.Assertions[idx] = ase

			if ase.Name == "body_contains" {
				ase.Result.Success = assert.ContainsArr(
					ioToString(resp.Body), ase.Expectations,
				)
				ase.Result.Msg = wra.addTestMessage(
					ase, "body did not contain %q", ase.Expectations,
				)
			}
			testcase.Assertions[idx] = ase

			if ase.Name == "header_key" {
				ase.Result = wra.assertHeader(resp.Header, ase)
				ase.Result.Msg = wra.addTestMessage(
					ase, "header key not %q", ase.Expectations,
				)
			}

			if ase.Name == "header_key_val" {
				ase.Result = wra.assertHeader(resp.Header, ase)
				ase.Result.Msg = wra.addTestMessage(
					ase, "header key val not %q", ase.Expectations,
				)
			}
			testcase.Assertions[idx] = ase
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

func (wra *tWebRA) assertHeader(header http.Header, ase tAssertion) (res tResult) {
	exp := make(map[string]string)
	for _, el := range ase.Expectations {
		arr := strings.Split(el, ":")
		if len(arr) < 2 {
			arr = append(arr, "")
		}
		arr[0] = strings.TrimSpace(arr[0])
		arr[1] = strings.TrimSpace(arr[1])
		exp[arr[0]] = arr[1]
	}
	res.Success = assert.IterHeader(header, exp)
	return
}

func (wra *tWebRA) makeTestCaseErrMessage(testcase tTestCase) (s string) {
	var arr []string
	for _, test := range testcase.Assertions {
		if test.Result.Msg != "" {
			arr = append(arr, test.Result.Msg)
		}
		s = strings.Join(arr, "; ")
	}
	return
}

func (wra *tWebRA) addTestMessage(ase tAssertion, msg string, itf ...interface{}) (s string) {
	if ase.Result.Success == false {
		if len(itf) == 1 {
			s = fmt.Sprintf(msg, itf[0])
		}
		if len(itf) == 2 {
			s = fmt.Sprintf(msg, itf[0], itf[1])
		}
	}
	return
}
