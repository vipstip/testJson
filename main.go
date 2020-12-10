package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fastjson"
	"io/ioutil"
	"log"
	"os"
	"testJson/StructJson"
	"time"
)

var byteValue []byte

func init() {
	jsonFile, _ := os.Open("rtb-request.json")
	byteValue, _ = ioutil.ReadAll(jsonFile)
}

func main() {

	//HalderEncodeJson(1000)
	//
	//EncodeFastJson()

	//HalderEasyJson(1000)

	//HalderJsonIter(1000)

	HalderFfjson(1000)

	//EncodeSimdjson()
}

func EncodeJson(channel chan StructJson.BidRequest) {

	var jsons StructJson.BidRequest

	json.Unmarshal(byteValue, &jsons)

	channel <- jsons
}

func HalderEncodeJson(total int) {
	defer Duration(Track("Encode-Json"))

	channel := make(chan StructJson.BidRequest, total)
	for i := 0; i < total; i++ {
		go EncodeJson(channel)
	}

	var datas = make([]StructJson.BidRequest, total)
	for i, _ := range datas {
		datas[i] = <-channel
	}

}

func EncodeFastJson() {
	defer Duration(Track("GetFastJson"))

	fmt.Println(fastjson.GetString(byteValue, "device", "ua"))
	fmt.Println(fastjson.GetInt(byteValue, "device", "height"))
	fmt.Println(fastjson.GetInt(byteValue, "device", "width"))
	fmt.Println(fastjson.GetInt(byteValue, "device", "dnt"))
	fmt.Println(fastjson.GetString(byteValue, "device", "language"))
	fmt.Println(fastjson.GetString(byteValue, "site", "id"))
	fmt.Println(fastjson.GetString(byteValue, "site", "page"))
	fmt.Println(fastjson.GetString(byteValue, "site", "referrer"))
	fmt.Println(fastjson.GetString(byteValue, "site", "hostname"))

}

func EncodeEasyJson(channel chan StructJson.BidRequest) {

	var d StructJson.BidRequest
	d.UnmarshalJSON(byteValue)
	channel <- d
}

func HalderEasyJson(total int) {
	defer Duration(Track("Easy-Json"))
	channel := make(chan StructJson.BidRequest, total)
	for i := 0; i < total; i++ {
		go EncodeEasyJson(channel)
	}

	var datas = make([]StructJson.BidRequest, total)
	for i, _ := range datas {
		datas[i] = <-channel
	}
}

func EncodeJsonIter(channel chan StructJson.BidRequest) {

	var data StructJson.BidRequest

	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	jsons.Unmarshal(byteValue, &data)

	channel <- data

}

func HalderJsonIter(total int) {
	defer Duration(Track("Json-Iter"))
	channel := make(chan StructJson.BidRequest, total)
	for i := 0; i < total; i++ {
		go EncodeJsonIter(channel)
	}

	var datas = make([]StructJson.BidRequest, total)
	for i, _ := range datas {
		datas[i] = <-channel
	}
}

func EncodeFfjson(channel chan StructJson.BidRequest) {

	var data StructJson.BidRequest

	ffjson.Unmarshal(byteValue, &data)

	channel <- data

}

func HalderFfjson(total int) {
	defer Duration(Track("Ffjson"))
	channel := make(chan StructJson.BidRequest, total)
	for i := 0; i < total; i++ {
		go EncodeFfjson(channel)
	}

	var datas = make([]StructJson.BidRequest, total)
	for i, _ := range datas {
		datas[i] = <-channel
	}

}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
