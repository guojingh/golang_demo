package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//    "description": "交易信息",
//    "type": " string",
//    "tr_type": {
//      "description": "交易类型",
//      "type": " string"
//    }

type trInfo struct {
	TrInfo struct {
		Description string `json:"description"`
		JsonType    string `json:"type"`
		TrType      struct {
			Description string `yaml:"description"`
			SonType     string `yaml:"sonType"`
		} `json:"tr_type"`
	} `json:"tr_info"`
}

func main() {

	file, err := os.Open("json/test.json")
	if err != nil {
		fmt.Println("读文件失败")
	}

	tr := &trInfo{}

	bytes := make([]byte, 1024)
	n, err := file.Read(bytes)
	err = json.Unmarshal(bytes[:n], tr)
	if err != nil {
		fmt.Println("结构体反序列化失败")
	}

	//fmt.Println(reflect.TypeOf(tr).Elem().)
	fmt.Printf("trInfo=%v", tr)
}
