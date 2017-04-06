package jsons

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type MorseJson struct{
	Words map[string]string
	Morse map[string]string
}

func (morseJson *MorseJson)ReadJsonData(strFileInfo string) {
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

