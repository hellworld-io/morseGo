package jsons

import (
	"testing"
	"path"
	"morseGo/constants"
)

func TestMorseJson_Convert(t *testing.T) {
	morseTestJson := new(MorseJson)
	dataFilePath := path.Join("../datas", constants.MorseDataFile)

	morseTestJson.ReadJsonData(dataFilePath)

	morseTestJson.Convert("abc d",".- -... -.-.  -..")
}
