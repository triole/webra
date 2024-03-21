package main

import (
	"os"
	"webra/src/logging"
)

var (
	lg logging.Logging
)

func main() {
	parseArgs()
	lg = logging.Init(CLI.LogLevel, CLI.LogFile, CLI.NoColors, CLI.JSONLog)

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
