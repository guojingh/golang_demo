package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func add(x, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	//解析参数
	b, _ := ioutil.ReadAll(r.Body)
	var param addParam
	json.Unmarshal(b, &param)
	fmt.Printf("client request x=%d,y=%d\n", param.X, param.Y)
	//业务逻辑
	ret := add(param.X, param.Y)
	fmt.Printf("ret=%d\n", ret)
	//返回响应
	result := addResult{
		Code: 200,
		Date: ret,
	}
	respBytes, _ := json.Marshal(result)
	n, err := w.Write(respBytes)
	if err != nil {
		fmt.Printf("err=%s", err)
	}
	fmt.Println(n)
}

func main() {
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
