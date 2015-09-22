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
	word := flag.String("word", "", "input your words")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	file, err := ioutil.ReadFile("./morseData.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &morseData)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(data)

	fmt.Println(morseData[*word])
}
