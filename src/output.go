package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var (
	red   = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
)

func (wra *tWebRA) report() {
	for _, testcase := range wra.TestSuite {
		if testcase.Result.Success == true {
			lg.Info(testcase.Name, logrus.Fields{
				"URL":  testcase.URL,
				"Type": "testcase",
			})
		} else {
			lg.Error(testcase.Name, logrus.Fields{
				"URL":  testcase.URL,
				"Msg":  testcase.Result.Msg,
				"Type": "testcase",
			})
		}
		for _, ase := range testcase.Assertions {
			lg.Debug(ase.Name, logrus.Fields{
				"URL":    testcase.URL,
				"Name":   ase.Name,
				"Exp":    ase.Expectations,
				"Result": ase.Result,
				"type":   "assertion",
			})
		}
	}
}
