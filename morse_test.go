package main

import (
	"testing"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
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

	//readJsonFile("./test.json")

	file, err := ioutil.ReadFile("./test.json")

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	//var morseStruct = new(MorseStruct)
	//morseUSData: make(map[string]interface{})
	var testMap map[string]interface{}
	err = json.Unmarshal(file, &testMap)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	//fmt.Println(MorseStruct.MorseUS.morseUSDatas["a"])
	fmt.Println(testMap["morseUS"])


}