package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("拨号失败", err)
	}
	defer conn.Close()
	fmt.Println("连接成功，开始聊天吧:")
	var input Interface
	input = &randPersonInformation{}
	for {
		personInformation, err := input.ReadPersonInformation()
		if err != nil {
			log.Println("获取用户输入失败,请重新录入:", err)
			continue
		}
		data, err := json.Marshal(personInformation)
		if err != nil {
			log.Println("无法编码个人信息:", err)
			continue
		}

		talk(conn, string(data))

	}
	//talk(conn, "你好")
	//talk(conn, "你是谁")
	//talk(conn, "今天天气怎样")
	//talk(conn, "再见")
}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败:", err)
	} else {
		data := make([]byte, 1024)
		vaildlen, err := conn.Read(data)
		if err != nil {
			log.Println("WARNING:读取服务器返回数据时出错:", err)
		} else {
			vaildData := data[:vaildlen]
			log.Println("去:", message, "回:", string(vaildData))
		}
	}
}
