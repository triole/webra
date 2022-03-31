package main

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

type tConf map[string]tConfEntry
type tConfEntry map[string]interface{}

// 	URLs              interface{} `toml:"url"`
// 	XStatusCodeEquals interface{} `toml:"x_status_code_equals"`
// 	XHeaderKeyVal     interface{} `toml:"x_header_key_val"`
// 	XBodyContains     interface{} `toml:"x_body_contains"`
// }

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

func (conf tConfEntry) hasKey(key string) bool {
	if _, ok := conf[key]; ok {
		return true
	}
	return false
}

func (conf tConfEntry) getKey(key string) (arr []string) {
	if conf.hasKey(key) == true {
		arr = interfaceToStrArr(conf[key])
	}
	return
}
