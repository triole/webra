package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/triole/logseal"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "WebRA"
	appDescription = "simple web request assertion tool"
	appMainversion = "0.1"
	lg             logseal.Logseal
)

var CLI struct {
	Config      string `help:"config toml file name, positional arg required" arg:"" optional:""`
	UserAgent   string `help:"user agent" default:"${userAgent}" short:"u"`
	Threads     int    `help:"max threads, default no of avail. cpu threads times 32" short:"n" default:"${threads}"`
	Timeout     int    `help:"request timeout in seconds" short:"t" default:"5"`
	Export      string `help:"export full test data into json file" short:"x"`
	LogFile     string `help:"log file" default:"${logfile}" short:"l"`
	LogLevel    string `help:"log level" short:"e" default:"info" enum:"trace,debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text"`
	LogJSON     bool   `help:"enable json log, instead of text one"`
	Debug       bool   `help:"debug mode" short:"d"`
	VersionFlag bool   `help:"display version" short:"V"`
}

func parseArgs() {
	curdir, _ := os.Getwd()
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
		kong.Vars{
			"logfile":   "stdout",
			"userAgent": appName + "/" + appMainversion + "." + getSubVersion(BUILDTAGS),
			"curdir":    curdir,
			"config":    path.Join(getBindir(), appName+".toml"),
			"threads":   strconv.Itoa(runtime.NumCPU() * 32),
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	if CLI.Config == "" {
		fmt.Printf("%s\n", "Error: Positional arg required. Please pass config toml file name.")
		os.Exit(1)
	}
	// ctx.FatalIfErrorf(err)
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func getSubVersion(buildtags string) (sv string) {
	t := rxFind(`_subversion: [0-9]+`, buildtags)
	arr := strings.Split(t, ":")
	if len(arr) > 1 {
		sv = strings.TrimSpace(arr[1])
	}
	return
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}

func getBindir() (s string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	s = filepath.Dir(ex)
	return
}
