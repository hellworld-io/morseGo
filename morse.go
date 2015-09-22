package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"flag"
)

var morseData map[string]interface{}

func main() {
	//code := flag.String("code", "", "input your morse code")
	//word := flag.String("word", "", "input your words")
	//var word []string
	//word = flag.Args()

	flag.Parse()

	/*
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	*/

	file, err := ioutil.ReadFile("./morseData.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
	}

	var outputString string

	for _, argument := range flag.Args() {
		outputString += argument + " "
	}

	fmt.Println(outputString)

	//fmt.Println(morseData[*word])
}
