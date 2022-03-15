package main

import (
	"encoding/json"
	"flag"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io"
	"learn/pkg/apis"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// handle error
	}
	rank := NewFatRateRank()
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("warning:建立连接失败:", err)
			continue
		}
		fmt.Println(conn)
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束连接:", conn)
	defer conn.Close()

	for {
		finalReceivedMessage := make([]byte, 0, 100)
		encounterError := false

		for {
			buf := make([]byte, 1024)
			vaild, err := conn.Read(buf)
			if err != nil {
				log.Println("WARNING,读取数据时出问题:", err)
				log.Println("重新读取", err)
				encounterError = true
				break
			}
			if vaild == 0 {
				break
			}
			vaildContect := buf[:vaild]
			finalReceivedMessage = append(finalReceivedMessage, vaildContect...)
			//fmt.Println("到这部了吗？")

			if encounterError {
				conn.Write([]byte("服务器获取数据失败,请重新输入"))
				continue
			}
			pi := apis.PersonInformation{}
			err = json.Unmarshal(finalReceivedMessage, &pi)
			if err != nil {
				conn.Write([]byte("输入数据无法解析"))
				continue
			}
			bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
			fatRate := gobmi.CalcFatRate(int(pi.Age), bmi, pi.Sex)
			rank.inputRecord(pi.Name, fatRate)
			fmt.Println(personFatRate)
			calRank, _ := getRand(pi.Name)
			//fmt.Printf("%s在排行榜中方的排名为%d,BMI为:%f,体脂率为%f\n", pi.Name, calRank, bmi, fatRate)
			resp := fmt.Sprintf("%s在排行榜中方的排名为%d,BMI为:%f,体脂率为%f\n", pi.Name, calRank, bmi, fatRate)
			data, err := json.Marshal(resp)
			if err != nil {
				log.Println("传回客户端信息反序列化失败:", err)
			}
			conn.Write(data)
			break

		}

	}
}
