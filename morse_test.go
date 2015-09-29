package main

import (
	"testing"
	"fmt"
)

type MorseStruct struct {
	MorseUS struct {
		morseUSData map[string]interface{}
	} `json:"morseUS"`
	MorseKOR struct {
		morseKORData map[string]interface{}
	} `json:"morseKOR"`
}

func TestParseToJsonData(t *testing.T){
	strFileName = "./test.json"
	readJsonFile(strFileName)

	//fmt.Println(MorseStruct.MorseUS.morseUSData["a"])
	//strMorseResult += morseData[strWord[idx:idx+1]].(string) + " "
	morseUSCode := morseData["morseUS"]["c"]

	if(morseUSCode == nil){
		fmt.Println("nil")
	}else{
		fmt.Println(morseUSCode)
	}
}