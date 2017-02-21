package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/tkyshm/goi18n-parser"
)

var (
	debug bool
	file  string
)

func main() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.StringVar(&file, "file", "", "file path")
	flag.Parse()

	a := goi18np.DefaultAnalyzer{
		Debug: debug,
	}

	rs := a.AnalyzeFromFile(file)
	out, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
