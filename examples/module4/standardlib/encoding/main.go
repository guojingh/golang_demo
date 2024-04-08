package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	data struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	gResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Results []data `json:"results"`
	}
)

func main() {

	uri := "http://127.0.0.1:8082/api/v1/community/?page=1&size=10"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println(resp.Body)
	/*	var gr gResponse
		json.Unmarshal(resp.Body., &gr)*/
}
