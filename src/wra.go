package main

import (
	"strings"
)

type tWebRA struct {
	TestSuite tTestSuite
}

type tTestSuite []tTestCase

func (arr tTestSuite) Len() int {
	return len(arr)
}

func (arr tTestSuite) Less(i, j int) bool {
	si := arr[i].Name + arr[i].URL
	sj := arr[j].Name + arr[j].URL
	return si < sj
}

func (arr tTestSuite) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type tTestCase struct {
	Name       string
	URL        string
	Assertions []tAssertion
	Result     tResult
}

type tAssertion struct {
	Name         string
	Expectations []string
	Result       tResult
}

type tResult struct {
	Success bool
	Msg     string
}

func initWebRA(conf tConf) (wra tWebRA) {
	for key, val := range conf {
		wra.TestSuite = append(wra.TestSuite, wra.initTestCases(key, val)...)
	}
	return
}

func (wra *tWebRA) initTestCases(name string, ce tConfEntry) (testcases []tTestCase) {
	for _, url := range interfaceToStrArr(ce["url"]) {
		tc := tTestCase{}
		tc.Name = name
		tc.URL = url

		tc = newAssertion(tc, ce, "connect", "success")
		tc = newAssertion(tc, ce, "x_status_code_equals", nil)
		tc = newAssertion(tc, ce, "x_header_key", nil)
		tc = newAssertion(tc, ce, "x_header_keyval", nil)
		tc = newAssertion(tc, ce, "x_body_contains", nil)

		testcases = append(testcases, tc)
	}
	return
}

func newAssertion(tc tTestCase, ce tConfEntry, key string, exp interface{}) tTestCase {
	var ase tAssertion
	ase.Name = strings.TrimPrefix(key, "x_")
	if exp != nil {
		ase.Expectations = interfaceToStrArr(exp)
	} else {
		if ce.hasKey(key) == true || key == "" {
			ase.Expectations = interfaceToStrArr(ce.getKey(key))
		}
	}
	if ase.Expectations != nil {
		tc.Assertions = append(tc.Assertions, ase)
	}
	return tc
}
