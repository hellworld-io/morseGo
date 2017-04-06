package flags

import (
	"fmt"
	"flag"
	"morseGo/constants"
	"os"
	"strings"
)

var (
	convertCmd = flag.NewFlagSet("convert",flag.ExitOnError)
	convertText = convertCmd.String("text","","Text to morse code ex) -text 'a b'")
	convertMorse = convertCmd.String("morse","","morse code to text ex) -morse '. .-  . .-'")

	TextArg string
	MorseArg string
)

func MorseMessage() {
	fmt.Println(constants.MorseUsage)
	flag.PrintDefaults()
}

func SetFlags() {
	flag.Usage = MorseMessage

	if len(os.Args) < 2 {
		MorseMessage()
		os.Exit(-1)
	}

	if os.Args[1] == "convert" {
		convertCmd.Parse(os.Args[2:])
	} else {
		MorseMessage()
		os.Exit(-1)
	}

	if convertCmd.Parsed() {
		if *convertText == "" && *convertMorse == "" {
			convertCmd.PrintDefaults()
			os.Exit(-1)
		}

		TextArg = strings.ToLower(*convertText)
		MorseArg = *convertMorse
	}
}
