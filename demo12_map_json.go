package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	UserName string   `json:"userName"`
	Age      int      `json:"age"`
	Hobby    []string `json:"hobby"`
}

func main() {
	fmt.Println("-----------------------------------")
	example()
	fmt.Println("-----------------------------------")
	example2()
}

func example() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = Result{"Tom", 30, []string{"读书", "运动"}}
	fmt.Println("map data :", res)

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error :", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json unmarshal error :")
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
}

func example2() {
	jsonStr := `{"number":1234566}`
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	//不确定类型使用interface来接收时，数字较大会被解析为科学计数法，需要强制类型转换
	fmt.Println(int(result["number"].(float64)))
}
