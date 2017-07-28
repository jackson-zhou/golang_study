package main

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"os"
)

func t1() {
	fmt.Println("Reading file")
	f, err := os.Open("post_data")
	if err != nil {
		fmt.Println("cannot read file")
		fmt.Println(err)
		return
		//错误
	}
	//txt,err := ioutil.ReadAll(f)
	var i_json interface{}
	if err := json.NewDecoder(f).Decode(&i_json); err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}

	imp_arr := i_json.(map[string]interface{})["imp"].([]interface{})
	for _, imp := range imp_arr {
		id := imp.(map[string]interface{})["id"].(string)
		fmt.Printf("id = %s\n", id)
		id = "abcdefg"
	}

	o_json, _ := json.Marshal(i_json)
	fmt.Println(string(o_json))
	fmt.Println("Hello, Go examples!")
}
