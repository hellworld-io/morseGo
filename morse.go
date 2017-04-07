package main

import (
	"morseGo/flags"
	"fmt"
	"morseGo/jsons"
	"morseGo/constants"
	"path"
)

func main() {
	morseJson := new(jsons.MorseJson)
	dataFilePath := path.Join("datas", constants.MorseDataFile)
	morseJson.ReadJsonData(dataFilePath)

	textArgs, morseArgs := flags.SetFlags()

	resultText, resultMorse := morseJson.Convert(textArgs, morseArgs)

	fmt.Printf("converted text : %s \nconverted morse code : %s\n",resultMorse, resultText)
}
