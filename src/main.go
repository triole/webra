package main

import (
	"os"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg = logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)

	conf := readConfigFile(CLI.Config)

	wra := initWebRA(conf)
	wra.processTestSuite()

	if CLI.Export != "" {
		writeJSONToFile(CLI.Export, wra)
	}

	wra.report()

	if !wra.isTestSuiteSuccessful() {
		os.Exit(1)
	}
}
