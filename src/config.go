package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type tConf map[string]tConfEntry
type tConfEntry map[string]interface{}

func readConfigFile(confFile string) (conf tConf) {
	var raw []byte
	if confFile != "" {
		var err error
		raw, err = os.ReadFile(confFile)
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
	if conf.hasKey(key) {
		arr = interfaceToStrArr(conf[key])
	}
	return
}

func (conf tConfEntry) getKeyStr(key string) (s string) {
	if conf.hasKey(key) {
		s = conf[key].(string)
	}
	return
}
