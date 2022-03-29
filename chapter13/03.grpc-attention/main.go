package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	pi := PersonInformation{
		Id:     0,
		Name:   "戴一杰",
		Sex:    "男",
		Tall:   1.75,
		Weight: 75,
		Age:    26,
	}
	jsonData, _ := json.Marshal(pi)
	fmt.Println((string(jsonData)))
	fmt.Println("========================")
	proData, _ := proto.Marshal(&pi)
	fmt.Println((string(proData)))
	fmt.Println("编码后========================")
	fmt.Println(base64.StdEncoding.EncodeToString(proData))
	fmt.Println("========================")

}
