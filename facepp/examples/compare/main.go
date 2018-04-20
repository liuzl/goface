package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/liuzl/goface/facepp"
	"log"
)

var (
	ak = flag.String("ak", "", "API Key")
	sk = flag.String("sk", "", "Secret Key")
)

func main() {
	flag.Parse()
	sdk, err := facepp.NewFaceSDK(*ak, *sk)
	if err != nil {
		log.Fatal(err)
	}
	compare, err := sdk.Compare()
	if err != nil {
		log.Fatal(err)
	}
	cr, body, err := compare.
		SetFace1("./demo-pic39.jpg", "image_file1").
		SetFace2("./demo-pic33.jpg", "image_file2").
		End()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)
	js, _ := json.Marshal(cr)
	fmt.Println(string(js))
}
