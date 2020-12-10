package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"os"
	"testJson/StructJson"

	//"testJson/StructJson"
	"time"
)

var byteValue1 []byte

func init() {
	jsonFile, _ := os.Open("rtb-request.json")
	byteValue1, _ = ioutil.ReadAll(jsonFile)
}

func main() {
	defer Duration2(Track2("JsonParse"))
	//t, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//t1, _,_,_ := jsonparser.Get(byteValue1, "imp", "mimes","[0]")
	//t2, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//t3, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//t4, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//t5, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//t6, _,_,_ := jsonparser.Get(byteValue1, "device", "geo", "lon")
	//fmt.Println(string(t))
	//fmt.Println(string(t1))
	//fmt.Println(string(t2))
	//fmt.Println(string(t3))
	//fmt.Println(string(t4))
	//fmt.Println(string(t5))
	//fmt.Println(string(t6))

	//err := jsonparser.ObjectEach(byteValue1, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	//	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	//	return nil
	//}, "imp","[0]")
	//if err != nil {
	//	panic(err)
	//}

	jsonparser.ArrayEach(byteValue1, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		js, _, _, _ := jsonparser.Get(value, "id")
		fmt.Println(string(js))
	}, "imp")

	bidID, _, _, _ := jsonparser.Get(byteValue1, "imp")
	var bid StructJson.BidRequest
	bid.ID, _ = jsonparser.GetString(byteValue1, "id")
	bid.ID, _ = jsonparser.GetString(byteValue1, "imp")
	abc, _, _, _ := jsonparser.Get(bidID, "[0]")
	fmt.Println(string(abc))
}

func Track2(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration2(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
