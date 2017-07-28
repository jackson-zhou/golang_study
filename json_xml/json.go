package main

import (
	_ "encoding/json"
	"encoding/xml"
	"fmt"
	_ "io/ioutil"
	"os"
)

func main() {
	fmt.Println("Reading file")
	f, _ := os.Open("adzone.91.xml.response")
	//txt,err := ioutil.ReadAll(f)
	var i_json interface{}
	if err := xml.NewDecoder(f).Decode(&i_json); err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}

	o_xml, err := xml.Marshal(i_json)
	fmt.Println("i_json", i_json)

	fmt.Println("err=", err)
	fmt.Println(string(o_xml))
}
