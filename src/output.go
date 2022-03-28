package main

import (
	"sort"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var (
	red   = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
)

func (wra *tWebRA) report() {
	sort.Sort(tTestSuite(wra.TestSuite))
	for _, testcase := range wra.TestSuite {
		if testcase.Result.Success == true {
			lg.Info(testcase.Name, logrus.Fields{
				"URL": testcase.URL,
			})
		} else {
			lg.Error(testcase.Name, logrus.Fields{
				"URL": testcase.URL,
				"Msg": testcase.Result.Msg,
			})
		}
	}
}
