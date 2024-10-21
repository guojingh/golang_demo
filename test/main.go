package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())

	json.Unmarshal()

}
