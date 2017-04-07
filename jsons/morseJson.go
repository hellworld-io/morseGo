package jsons

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"strings"
)

type MorseJson struct{
	Words map[string]string
	Morse map[string]string
}

func (morseJson *MorseJson) ReadJsonData(strFileInfo string) {
	file, err := ioutil.ReadFile(strFileInfo)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	err = json.Unmarshal(file, morseJson)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func (morseJson *MorseJson) Convert(textArg string, morseArg string) (convertedText string, convertedMorse string) {
	for _, texts := range textArg {
		if string(texts) != " " {
			convertedMorse += morseJson.Words[string(texts)] + " "
		} else {
			convertedMorse += morseJson.Words[string(texts)]
		}
	}

	eachMorseCode := strings.Split(morseArg," ")
	for _, morsecodes := range eachMorseCode {
		if morsecodes == "" {
			convertedText += " "
		}
		convertedText += morseJson.Morse[string(morsecodes)]
	}

	return convertedText, convertedMorse
}

