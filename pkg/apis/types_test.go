package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"

	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

//func TestMarshalJson(t *testing.T) {
//
//	personInformation := PersonInformation{
//		Name:   `"戴一杰"`,
//		Sex:    "男",
//		Tall:   1.75,
//		Weight: 75,
//		Age:    26,
//	}
//
//	fmt.Printf("%v\n", personInformation)
//
//	data, err := json.Marshal(personInformation)
//	if err != nil {
//		log.Fatal(err)
//
//	}
//	fmt.Println("Marshal的结果是(原生):", data)
//	fmt.Println("Marshal的结果是(string):", string(data))
//
//}
func TestUnmarshalJson(t *testing.T) {
	data := `{"name":"戴一节","sex":"男","tall":1.76,"weight":75,"age":26}`
	personInformation := PersonInformation{}
	json.Unmarshal([]byte(data), &personInformation)
	fmt.Println(personInformation)
}

func TestMarshalYaml(t *testing.T) {

	personInformation := PersonInformation{
		Name:   "戴一杰",
		Sex:    "男",
		Tall:   1.75,
		Weight: 75,
		Age:    26,
	}

	fmt.Printf("%v\n", personInformation)

	data, err := json.Marshal(personInformation)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Marshal的结果是(原生):", data)
	fmt.Println("Marshal的结果是(string):", string(data))

}
func TestUnmarshalYaml(t *testing.T) {
	data := ` {"name":"戴一节","sex":"男","tall":1.76,"weight":75,"age":26}`
	personInformation := PersonInformation{}
	yaml.Unmarshal([]byte(data), &personInformation)
	fmt.Println(personInformation)
}

//func TestMarshalProtobuf(t *testing.T) {
//	personalInformation := &PersonInformation{
//		Name:   `"小"强""`,
//		Sex:    "男",
//		Tall:   1.70,
//		Weight: 71,
//		Age:    35,
//	}
//	data, err := proto.Marshal(personalInformation)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(data)
//	fmt.Println(string(data))
//	// 通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行保存，必须转码。
//	// 选择的通用转码是：base64
//	output64Data := base64.StdEncoding.EncodeToString(data)
//	fmt.Println(">>>>>", output64Data)
//}
func TestMarshalProtobuf(t *testing.T) {
	personalInformation := &PersonInformation{
		Name:   `"小"强""`,
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}
	data, err := proto.Marshal(personalInformation)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
	// 通常在非程序交互过程中，要保留原生protobuf，可以直接写入文件。如果想要单行保存，必须转码。
	// 选择的通用转码是：base64
	output64Data := base64.StdEncoding.EncodeToString(data)
	fmt.Println(">>>>>", output64Data)
}
