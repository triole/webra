package main

import "webra/src/logging"

var (
	lg logging.Logging
)

func main() {
	parseArgs()
	lg = logging.Init(CLI.LogFile, CLI.JSONLog)

	wra := readConfigFile(CLI.Config)
	wra.runTests()
	wra.report()
}
