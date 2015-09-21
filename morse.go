package main

import (
	"fmt"
	"os"
)


type morseCode struct {
	code string
	text string
}

func convertToMorseByString(strInput string) string{

	if strInput != "" {
		return "morse code"
	}

	return "mistake input string"
}

func convertToStringByMorse(strInput string) string{


	return "mistake input string"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s String or Morse Code\n", os.Args[0])
		os.Exit(-1)
	}

	strInput := os.Args[1]

	fmt.Println(convertToMorseByString(strInput))

}
