package main

import (
	"strconv"
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
	Name   string
	URL    string
	Tests  []tTest
	Result tResult
}

type tTest struct {
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
		for _, url := range makeStringList(val.URLs) {
			var tc tTestCase
			tc.Name = key
			tc.URL = url

			var test tTest
			test.Name = "Connect"
			// TODO: think about a meaningful and functional expectation
			test.Expectations = makeExpectations("success")
			tc.Tests = appendExpectations(tc.Tests, test)

			if val.XStatusCodeEquals != nil {
				var test tTest
				test.Name = "StatusCodeEquals"
				test.Expectations = makeExpectations(val.XStatusCodeEquals)
				tc.Tests = appendExpectations(tc.Tests, test)
			}

			if val.XHeaderKeyVal != nil {
				var test tTest
				test.Name = "HeaderKeyVal"
				test.Expectations = makeExpectations(val.XHeaderKeyVal)
				tc.Tests = appendExpectations(tc.Tests, test)
			}

			if val.XBodyContains != nil {
				var test tTest
				test.Name = "BodyContains"
				test.Expectations = makeExpectations(val.XBodyContains)
				tc.Tests = appendExpectations(tc.Tests, test)
			}

			wra.TestSuite = append(wra.TestSuite, tc)
		}
	}
	return
}

func appendExpectations(tests []tTest, test tTest) []tTest {
	if len(test.Expectations) > 0 {
		tests = append(tests, test)
	}
	return tests
}

func makeExpectations(itf interface{}) (exp []string) {
	switch val := itf.(type) {
	case int:
		exp = []string{strconv.Itoa(val)}
	case int64:
		exp = []string{strconv.Itoa(int(val))}
	case string:
		exp = []string{val}
	case []interface{}:
		exp = makeStringList(val)
	}
	return
}

func makeStringList(itf interface{}) (arr []string) {
	switch val := itf.(type) {
	case string:
		arr = []string{val}
	case []string:
		arr = val
	case []interface{}:
		for _, el := range val {
			arr = append(arr, el.(string))
		}
	}
	return
}
