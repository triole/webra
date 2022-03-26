package main

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

type tWebRA struct {
	TestsIterator tTestsIterator
	TestsMap      tTestsMap
	ResultsMap    tResultsMap
}

type tTestsIterator []string
type tTestsMap map[string]tTest
type tResultsMap map[string]tResult

type tTest struct {
	URL                 string   `toml:"url"`
	ExpStatusCode       int      `toml:"exp_status_code"`
	ExpHeaderSubstrings []string `toml:"exp_header_substring"`
	ExpBodySubstrings   []string `toml:"exp_body_substring"`
	Result              tResults
}

type tResults struct {
	Success bool
	Msg     string
	List    []tResult
}

type tResult struct {
	TestName string
	Success  bool
	Got      interface{}
	Exp      interface{}
	Desc     string
	Msg      string
}

func readConfigFile(confFile string) (conf tWebRA) {
	treq := make(tTestsMap)
	var raw []byte
	if confFile != "" {
		var err error
		raw, err = ioutil.ReadFile(confFile)
		if err != nil {
			log.Fatalf("Error reading config %q, %q", confFile, err)
		}
		err = toml.Unmarshal(raw, &treq)
		if err != nil {
			log.Fatalf("Error unmarshal %q, %q", confFile, err)
		}
	}
	conf.TestsMap = treq
	conf.TestsIterator = makeRequestIterator(raw)
	return
}

func makeRequestIterator(raw []byte) (iter []string) {
	arr := rxFindAll(`(^|\n)\[[a-zA-Z0-9-_]+\](\n)`, string(raw))
	for _, el := range arr {
		m := rxFind(`[a-zA-Z0-9-_]+`, el)
		if m != "" {
			iter = append(iter, m)
		}
	}
	return
}
