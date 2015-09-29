package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
	"os"
	"strings"
)

var morseData map[string] map[string] interface{}

var strFileName string

var bAlphabetToMorse = flag.Bool("atm", false, "To need Alphabet words ex) -atm ab cd")
//var bKorToMorse = flag.Bool("ktm",false, "To need Kor words")
var bMorseToAlphabet = flag.Bool("mta", false, "To need morse codes for alphabet ex) -mta '. .-  . .-' ")
//var bMorseToKor = flag.Bool("mtk", false, "To need morse codes for Kor")

func makeToStringByArgsAlphabet(strArgs []string) string {
	var strArgsResult string

	if(*bAlphabetToMorse == true){
		for _, argument := range strArgs {
			strArgsResult += argument + " "
		}

		strArgsResult = strings.ToLower(strArgsResult)
	}else if(*bMorseToAlphabet == true){
		for _, argument := range strArgs {
			strArgsResult += argument
		}
	}else{
		log.Fatal("Arguments error !!! no flag")
		os.Exit(-1)
	}

	if (len(strArgs) == 0){
		log.Fatal("Arguments is null !!!")
		os.Exit(-1)
	}

	return strArgsResult
}

func convertToByArgs(strWord string) string{
	var strMorseResult string

	if(*bAlphabetToMorse == true) {
		for idx := 0; idx < len(strWord); idx++ {
			if(morseData["wordsMorseUS"][strWord[idx:idx+1]] == nil){
				log.Fatal("There are no matching words.")
				os.Exit(-1)
			}
			strMorseResult += morseData["wordsMorseUS"][strWord[idx:idx+1]].(string) + " "
		}
	}else if(*bMorseToAlphabet == true){
		arrMorseWord := strings.Split(strWord," ")
		for idx := 0; idx < len(arrMorseWord); idx++ {
			if(arrMorseWord[idx] == ""){
				arrMorseWord[idx] = " "
			}

			if(morseData["morseWordsUS"][arrMorseWord[idx]] == nil){
				log.Fatal("There are no matching words.")
				os.Exit(-1)
			}
			strMorseResult += morseData["morseWordsUS"][arrMorseWord[idx]].(string)
		}
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

	var strArgs string

	strArgs = makeToStringByArgsAlphabet(flag.Args())

	strFileName = "morseData.json"
	readJsonFile(strFileName)

	var strMorse = convertToByArgs(strArgs)

	fmt.Println(strMorse)
}
