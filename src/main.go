package main

import "webra/src/logging"

var (
	lg logging.Logging
)

func main() {
	parseArgs()
	lg = logging.Init(CLI.LogFile, CLI.JSONLog)

	conf := readConfigFile(CLI.Config)

	wra := initWebRA(conf)
	wra.runTestSuite()
	wra.report()
}
