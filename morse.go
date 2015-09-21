package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)


type morseCode struct {
	code string
	text string
}

type YourJson struct {
	YourSample []struct {
		data map[string]string
	}
}

func convertToMorseByString(strInput string) string{

	// Alphabet to Morse Code
	if strInput != "" {
		return "morse code"
	}

	return "mistake input string"
}

func convertToStringByMorse(strInput string) string{
	// Morse Code to String

	return "mistake input string"
}

func main() {
	/*
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s String or Morse Code\n", os.Args[0])
		os.Exit(-1)
	}

	strInput := os.Args[1]

	fmt.Println(convertToMorseByString(strInput))
	*/

	var data map[string]interface{}
	file, err := ioutil.ReadFile("./morseData.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	strInput := os.Args[1]
	fmt.Println(data[strInput])

	/*
	file, e := ioutil.ReadFile("./test.json")

	if (e != nil) {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}


	fmt.Println(string(file))

	var yourJson YourJson
	json.Unmarshal(file, &yourJson)
	fmt.Printf("Results: %v\n", yourJson)

	dec := json.NewDecoder(bytes.NewReader(file))
	var d YourJson
	fmt.Println(dec.Decode(&d))
	*/
}
