package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
)

var morseData map[string]interface{}

func makeToStringArrayByArgs(strArgs []string) string {
	var strArgsResult string

	for _, argument := range strArgs {
		strArgsResult += argument + " "
	}
	return strArgsResult
}

func convertToMorseByWord(strWord string) string{
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
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	readJsonFile()
	var strArgs = makeToStringArrayByArgs(flag.Args())
	var strMorse = convertToMorseByWord(strArgs)

	fmt.Println(strMorse)
}
