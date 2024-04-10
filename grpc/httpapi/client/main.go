package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Date int `json:"date"`
}

func main() {
	//通过 HTTP 调用远程 add 服务
	url := "http://127.0.0.1:9090/add"
	param := addParam{
		X: 20,
		Y: 10,
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)
	fmt.Printf("resp=%v", respData)
	fmt.Println(respData.Date)
}