package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learn/pkg/apis"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	w, err := c.WatchPersons(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("启动Watcher失败:", err)
	}
	for true {
		pi, err := w.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("服务器告知说完了")
				break
			}
			log.Fatal("接受异常:", err)
		}
		log.Println("收到变动:", pi.String())

	}

	ret, err := c.Register(context.TODO(), &apis.PersonInformation{Name: "tom"})
	if err != nil {
		log.Fatal("注册失败:", err)
	}
	log.Println("注册成功", ret)
	log.Println("开始连续注册")
	regClient, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败:", err)
	}

	resp, err := regClient.CloseAndRecv()
	if err != nil {
		log.Fatal("接收失败:", err)
	}
	log.Println("批量注册结果:", resp.String())
}
