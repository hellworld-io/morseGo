package main

import (
	"testing"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

func TestParseToJsonData(t *testing.T){

	file, err := ioutil.ReadFile("./test.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(morseData)

}

func TestConvertToMorseCodeByJsonData(t *testing.T){

}

func TestConvertToWordsByJsonData(t *testing.T){

}
