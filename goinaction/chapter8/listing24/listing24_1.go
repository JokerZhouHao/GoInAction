package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	InsertKeyHere string `json:"insert-key-here"`
	Key           string `json:"key"`
}

func main() {
	uri := "http://echo.jsontest.com/insert-key-here/insert-value-here/key/value"

	// 请求
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()

	// 解析json
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(result)

	// 将结构体转化为json二进制串
	pretty, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(pretty)
	fmt.Println(string(pretty))

	// 将json串转化为map
	var c map[string]interface{}
	err = json.Unmarshal(pretty, &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c["key"])

}
