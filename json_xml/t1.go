package main

import (
	_ "encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Reading file")
	f, err := os.Open("adzone.91.xml.response")
	txt, err := ioutil.ReadAll(f)
	fmt.Println("txt:", string(txt))
	var v interface{}
	err = xml.Unmarshal(txt, v)
	fmt.Println("i_xml", v)
	fmt.Println("err=", err)
	o_xml, err := xml.Marshal(v)
	fmt.Println(string(o_xml))
}
