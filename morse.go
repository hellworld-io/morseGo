package main

import (
	"morseGo/flags"
)

/*
func makeToStringByArgsAlphabet(strArgs []string) string {
	var strArgsResult string



	fmt.Println("args =======> ",strArgs)
	for _, argument := range strArgs {
		fmt.Println("argument =======> ",argument)
		strArgsResult += argument
	}

	fmt.Println("strArgsResult =======> ",strArgsResult)

	if *bAlphabetToMorse != "" {
		strArgsResult = strings.ToLower(strArgsResult)
	}

	return strArgsResult
}

func convertToByArgs(strWord string) string {
	var strMorseResult string

	if *bAlphabetToMorse != "" {
		for idx := 0; idx < len(strWord); idx++ {
			if morseData["wordsMorseUS"][strWord[idx:idx+1]] == nil {
				log.Fatal("There are no matching words.")
				os.Exit(-1)
			}
			strMorseResult += morseData["wordsMorseUS"][strWord[idx:idx+1]].(string) + " "
		}
	} else if *bMorseToAlphabet != "" {
		arrMorseWord := strings.Split(strWord, " ")
		for idx := 0; idx < len(arrMorseWord); idx++ {
			if arrMorseWord[idx] == "" {
				arrMorseWord[idx] = " "
			}

			if morseData["morseWordsUS"][arrMorseWord[idx]] == nil {
				log.Fatal("There are no matching morse codes.")
				os.Exit(-1)
			}
			strMorseResult += morseData["morseWordsUS"][arrMorseWord[idx]].(string)
		}
	}

	return strMorseResult
}

*/
func main() {

	flags.SetFlags()

	//strFileName = "./morseData.json"
	//morseJsonData := new(commons.MorseObject)
	//morseJsonData.ReadJsonData(strFileName)

	//fmt.Println(morseJsonData.Words["1"])

	//flag.Usage = flags.MorseMessage
	//
	//flag.Parse()

	//var strArgs string

	//checkFlags()

	//fmt.Println(*bAlphabetToMorse)
	//strArgs = makeToStringByArgsAlphabet(flag.Args())
	//
	//morseJsonData := new(jsons.MorseJson)
	//morseJsonData.ReadJsonData(constants.MorseDataFile)
	//fmt.Println(morseJsonData.Words["1"])
	//
	//var strMorse = convertToByArgs(strArgs)
	//
	//fmt.Println(strMorse)

}
