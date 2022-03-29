package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonData := "{\"Name\":\"戴一杰\",\"Sex\":\"男\",\"Tall\":1.75,\"Weight\":75,\"Age\":26}"
	proDataBase64 := "EgnmiLTkuIDmnbAaA+eUtyUAAOA/LQAAlkIwGg=="
	fmt.Printf("jsonData：%s", jsonData)
	pi1 := &PersonInformation{}
	json.Unmarshal([]byte(jsonData), &pi1)
	fmt.Println("解析json：")
	fmt.Printf("%+v\n", string(pi1.Sex))
	fmt.Printf("%s\n", pi1)

	//fmt.Println("map=======================")
	//test := map[string]string{"name": "一杰"}
	//pi3 := make(map[string]string)
	//json.Unmarshal([]byte(test), &pi3)
	//fmt.Println("解析json：")
	//
	//fmt.Printf("%+v\n", pi3)

	fmt.Println("protobuf=======================")
	pi2 := &PersonInformation{}
	proData, _ := base64.StdEncoding.DecodeString(proDataBase64)
	proto.Unmarshal([]byte(proData), pi2)
	fmt.Println("解析protobuf：")
	fmt.Printf("%s\n", pi2)
}
