package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Tl struct {
	Rolltype int
	Timeout  int
}

type Dsp struct {
	Participation_bid    bool //是否参与竞价
	Use_cdn              int  //使用乐视CDN来转码和承压，0表示不用，1表示用。 Cookie_mapping_url   string
	Qps                  int
	Timeoutlimit         []Tl
	Transform_type       int
	Token                string
	Bid_url              string
	Win_notice_encrypted bool
	Multi_deal           bool
	Win_notice_url       string
	Show_ad              int
	Use_cookie_mapping   bool
	Auction_type         int
	Icon_url             string `json:"icon_url,omitempty"`
}

func t1() {
	fmt.Println("Reading file")
	f, _ := os.Open("dsp.1")
	txt, _ := ioutil.ReadAll(f)
	fmt.Println("txt=", string(txt))
	fmt.Println("f=", f)
	var dsp Dsp
	decoder := json.NewDecoder(f)
	fmt.Println("decoder=", decoder)
	err := decoder.Decode(&dsp)
	err = json.Unmarshal(txt, &dsp)
	fmt.Println("error=", err)
	fmt.Println("dsp=", dsp)
	fmt.Println("dsp.Bid_url=", dsp.Bid_url)
	fmt.Println("===============================")
	xml_txt, err := xml.Marshal(&dsp)
	fmt.Println("error=", err)
	fmt.Println("xml=", string(xml_txt))
}
func main() {
	t1()
}
