package main

import (
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	str := `{
			"alipay_trade_create_response": {
				"code": "10000",
				"msg": "Success",
				"out_trade_no": "20150423001001",
				"trade_no": "2015042321001004720200028594"
			},
			"sign": "ERITJKEIJKJHKKKKKKKHJEREEEEEEEEEEE"
	}`

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(str), &body); err != nil {
		log.Error(err)
	}

	fmt.Printf("%s\n", body["alipay_trade_create_response"].(map[string]interface{})["code"])
	fmt.Println(body["sign"])
}
