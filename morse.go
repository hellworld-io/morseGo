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

var strFileName string

var fMorseUS = flag.Bool("um", false, "To need US words")
var fMorseKOR = flag.Bool("km",false, "To need KOR words")
var fMorseCode = flag.Bool("mw", false, "To need morse codes")

func makeToStringByArgs(strArgs []string) string {
	var strArgsResult string

	if (len(strArgs) == 0){
		log.Fatal("Args is null !!!")
		os.Exit(-1)
	}

	for _, argument := range strArgs {
		strArgsResult += argument + " "
	}
	return strArgsResult
}

func convertToMorseByArgs(strWord string) string{
	var strMorseResult string

	for idx := 0; idx < len(strWord); idx++ {
		if(morseData[strWord[idx:idx+1]] == nil){
			log.Fatal("There are no matching words.")
			os.Exit(-1)
		}
		strMorseResult += morseData[strWord[idx:idx+1]].(string) + " "
	}

	return strMorseResult
}

func readJsonFile(strFileName string) {
	file, err := ioutil.ReadFile("./" + strFileName)

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

	fmt.Println(*fMorseUS)
	var strArgs = makeToStringByArgs(flag.Args())

	strFileName = "morseData.json"
	readJsonFile(strFileName)

	var strMorse = convertToMorseByArgs(strArgs)

	fmt.Println(strMorse)
}
