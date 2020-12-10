package main

import (
	"fmt"
	"github.com/minio/simdjson-go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func printKey(iter simdjson.Iter, key string) (err error) {

	obj, tmp, elem := &simdjson.Object{}, &simdjson.Iter{}, simdjson.Element{}

	for {
		typ := iter.Advance()

		switch typ {
		case simdjson.TypeRoot:
			if typ, tmp, err = iter.Root(tmp); err != nil {
				return
			}

			if typ == simdjson.TypeObject {
				if obj, err = tmp.Object(obj); err != nil {
					return
				}
				e := obj.FindKey(key, &elem)
				fmt.Println(elem.Iter.Object(obj))
				if e != nil && elem.Type == simdjson.TypeString {
					v, _ := elem.Iter.StringBytes()
					fmt.Println(string(v))
				}
			}

		default:
			return
		}
	}
}

func main() {
	defer Duration1(Track1("Simdjson"))
	jsonFile, _ := os.Open("bid-request.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	parsed, err := simdjson.ParseND(byteValue, nil)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	//iter := parsed.Iter()
	//_, _, _ = iter.Root(iter)
	fmt.Println()
	printKey(parsed.Iter(), "device")
}

func Track1(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration1(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
