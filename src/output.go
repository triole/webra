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
	for _, el := range wra.TestsIterator {
		test := wra.TestsMap[el]
		if test.Result.Success == true {
			lg.Info(el, logrus.Fields{
				"URL": test.URL,
			})
		} else {
			lg.Error(el, logrus.Fields{
				"Msg": test.Result.Msg,
				"URL": test.URL,
			})
		}
	}
}
