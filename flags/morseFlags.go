package flags

import (
	"flag"
	"fmt"
	"morseGo/constants"
	"os"
	"strings"
)

var (
	convertCmd   = flag.NewFlagSet("convert", flag.ExitOnError)
	convertText  = convertCmd.String("text", "", "Text to morse code ex) -text 'a b'")
	convertMorse = convertCmd.String("morse", "", "morse code to text ex) -morse '. .-  . .-'")
)

func defaultMessage() {
	fmt.Println(constants.MorseUsage)
	flag.PrintDefaults()
}

func SetFlags() (textArgs string, morseArg string) {
	flag.Usage = defaultMessage

	if len(os.Args) < 2 {
		defaultMessage()
		os.Exit(-1)
	}

	if os.Args[1] == "convert" {
		convertCmd.Parse(os.Args[2:])
	} else {
		defaultMessage()
		os.Exit(-1)
	}

	if convertCmd.Parsed() {
		if *convertText == "" && *convertMorse == "" {
			convertCmd.PrintDefaults()
			os.Exit(-1)
		}

		textArgs = strings.ToLower(*convertText)
		morseArg = *convertMorse
	}
	return textArgs, morseArg
}
