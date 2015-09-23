package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
)

var morseData map[string]interface{}

func main() {
	flag.Parse()

	file, err := ioutil.ReadFile("./morseData.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
	}

	var outputString string
	var outMorse string

	for _, argument := range flag.Args() {
		outputString += argument + " "
	}

	for idx := 0; idx < len(outputString); idx++ {
		outMorse += morseData[outputString[idx:idx+1]].(string) + " "
	}

	fmt.Println(outMorse)
}
