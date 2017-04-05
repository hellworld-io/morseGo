package useage

import (
	"fmt"
	"flag"
	"morseGo/commons/constants"
)

func MorseUseage() {
	fmt.Println(constants.MorseUsage)
	flag.PrintDefaults()
}
