package jsonUtils

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"

)

type MorseObject struct{
	Words map[string]string
	Morse map[string]string
}

func ReadJsonData(strFileInfo string) {
	file, err := ioutil.ReadFile(strFileInfo)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	var jsonData MorseObject

	err = json.Unmarshal(file, &jsonData)

	//fmt.Println(jsonData.Words)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

