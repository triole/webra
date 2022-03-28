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
	wra.processTestSuite()

	if CLI.Export != "" {
		writeJSONToFile(CLI.Export, wra)
	}

	wra.report()
}
