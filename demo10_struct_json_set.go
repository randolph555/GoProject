package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `josn:"msg"`
}

func main() {
	res := Result{Code: 200, Message: "success"}

	toJson(&res)
	setData(&res)
	toJson(&res)
}

func toJson(res *Result) {
	json, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("json data :", string(json))
}

func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}
