package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
	"os"
)

var morseData map[string]interface{}


func makeToStringByArgs(strArgs []string) string {
	var strArgsResult string

	for _, argument := range strArgs {
		strArgsResult += argument + " "
	}
	return strArgsResult
}

func convertToMorseByArgs(strWord string) string{
	var strMorseResult string

	for idx := 0; idx < len(strWord); idx++ {
		strMorseResult += morseData[strWord[idx:idx+1]].(string) + " "
	}

	return strMorseResult
}

func readJsonFile() {
	file, err := ioutil.ReadFile("./morseData.json")

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func main() {
	flag.Parse()

	readJsonFile()
	var strArgs = makeToStringByArgs(flag.Args())
	var strMorse = convertToMorseByArgs(strArgs)

	fmt.Println(strMorse)
}
