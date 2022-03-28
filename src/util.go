package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func printErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func rxFindAll(rx string, str string) (arr []string) {
	re := regexp.MustCompile(rx)
	arr = re.FindAllString(str, -1)
	return
}

func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}

func rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func ioToString(io io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(io)
	return buf.String()
}

func writeJSONToFile(filename string, data tWebRA) {
	rankingsJSON, _ := json.Marshal(data)
	err := ioutil.WriteFile(filename, rankingsJSON, 0644)
	if err != nil {
		fmt.Printf("error saving json file %q\n", filename)
	}
}
