package main

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

type tConf map[string]tConfEntry

type tConfEntry struct {
	URL               string      `toml:"url"`
	XStatusCodeEquals interface{} `toml:"x_status_code_equals"`
	XHeaderKeyVal     interface{} `toml:"x_header_key_val"`
	XBodyContains     interface{} `toml:"x_body_contains"`
}

func readConfigFile(confFile string) (conf tConf) {
	var raw []byte
	if confFile != "" {
		var err error
		raw, err = ioutil.ReadFile(confFile)
		if err != nil {
			log.Fatalf("Error reading config %q, %q", confFile, err)
		}
		err = toml.Unmarshal(raw, &conf)
		if err != nil {
			log.Fatalf("Error unmarshal %q, %q", confFile, err)
		}
	}
	return
}
