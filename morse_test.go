package main
import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"fmt"
)

type YourJson struct {
	YourSample []struct {
		data map[string]string
	}
}

func TestParseToJsonData(t *testing.T){
	c, _ := ioutil.ReadFile("test")

	dec := json.NewDecoder(bytes.NewReader(c))
	var d YourJson
	fmt.Println(dec.Decode(&d))
}
