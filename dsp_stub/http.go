package main

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Receive reqeust")
	var i_json interface{}
	if err := json.NewDecoder(req.Body).Decode(&i_json); err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	i_json_obj := i_json.(map[string]interface{})
	imp_arr := i_json_obj["imp"].([]interface{})

	var adzoneid float64
	for _, imp := range imp_arr {
		adzoneid = imp.(map[string]interface{})["adzoneid"].(float64)
		fmt.Printf("adzoneid = %.0f\n", adzoneid)
	}
	impid := imp_arr[0].(map[string]interface{})["id"].(string)
	fmt.Println("===================")

	filename := "adzone." + strconv.FormatFloat(adzoneid, 'f', -1, 32)
	f, err := os.Open(filename)
	if err != nil {
		err_msg := string("cannot read file:") + filename
		w.Write([]byte(err_msg))
		fmt.Println(err_msg)
		return
	}

	var o_json interface{}
	json.NewDecoder(f).Decode(&o_json)
	o_json_obj := o_json.(map[string]interface{})
	o_json_obj["id"] = i_json_obj["id"]

	/*
		o_json_obj["seatbid"].([]interface{})[0]\
			.(map[string] interface{})["bid"].([]interface{})[0]\
			.(map[string] interface{})["impid"] = impid
	*/
	o_json_obj["seatbid"].([]interface{})[0].(map[string]interface{})["bid"].([]interface{})[0].(map[string]interface{})["impid"] = impid

	o_json_txt, _ := json.Marshal(o_json)
	fmt.Printf("%s\n", o_json_txt)
	w.Write([]byte(o_json_txt))
}
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("dsp_stub port")
		return
	}

	http.HandleFunc("/", ProcessRequest)
	HTTP_PORT := ":380" + args[1]
	fmt.Println("Listen on: " + HTTP_PORT[1:])
	http.ListenAndServe(HTTP_PORT, nil)
}
