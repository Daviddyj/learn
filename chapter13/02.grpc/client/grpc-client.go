package main

import (
	"context"
	"google.golang.org/grpc"
	"learn/pkg/apis"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonInformation{
		Id:     0,
		Name:   "dyj",
		Sex:    "男",
		Tall:   1.75,
		Weight: 75,
		Age:    26,
	})
	if err != nil {
		log.Fatal("注册失败:", err)
	}
	log.Println("注册成功", ret)
	log.Println("开始连续注册")
	regClient, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败:", err)
	}

	if err := regClient.Send(&apis.PersonInformation{
		Id:     1,
		Name:   "daidai",
		Sex:    "男",
		Tall:   1.86,
		Weight: 75,
		Age:    27,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regClient.Send(&apis.PersonInformation{
		Id:     2,
		Name:   "benben",
		Sex:    "男",
		Tall:   1.56,
		Weight: 75,
		Age:    27,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regClient.Send(&apis.PersonInformation{
		Id:     3,
		Name:   "lulu",
		Sex:    "男",
		Tall:   1.86,
		Weight: 75,
		Age:    27,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regClient.Send(&apis.PersonInformation{
		Id:     4,
		Name:   "mengmeng",
		Sex:    "男",
		Tall:   1.76,
		Weight: 75,
		Age:    27,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	resp, err := regClient.CloseAndRecv()
	if err != nil {
		log.Fatal("接收失败:", err)
	}
	log.Println("批量注册结果:", resp.String())

	//

	top, err := c.GetTop(context.TODO(), &apis.Null{})
	if err != nil {
		log.Println("获取体脂排行榜出错:", err)
	}
	log.Println("体脂排行榜为:", top)

}
