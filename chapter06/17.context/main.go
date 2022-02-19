package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	withCancel()
}

//func withTimeout() {
//	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
//}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消:distributeMainFrame")
		return
	default:
	}
	fmt.Println("部署:distributeMainFrame")

}
func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消:distributeMainBody")
		return
	default:
	}
	fmt.Println("部署:distributeMainBody")

}
func distributeCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消:distributeMainBody")
		return
	default:
	}
	fmt.Println("部署:distributeMainBody")

}

func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	go buyNoodles(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了，取消购买所有食材")
	cancel() //当调用cancel后，所以衍生出来的context都会cancel
	time.Sleep(3 * time.Second)

}
func buyNoodles(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): //todo 介绍一下 struct{}    //struct{}为空对象，是为了节省空间
		fmt.Println("收到消息，不买面了")
		return
	default:

	}
	fmt.Println("买面")
}
func buyOil(ctx context.Context) {

	fmt.Println("去买油")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default:

	}
	fmt.Println("买油")

}
func buyEgg(ctx1 context.Context) {
	ctx, _ := context.WithCancel(ctx1)
	//defer cancel()      //无论是否调用衍生出来的ctx的cancel，Done返回的Channel都会关闭
	fmt.Println("去买蛋")
	//time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:

	}
	fmt.Println("买蛋")
	go buySEgg(ctx)
	go buyBEgg(ctx)
}
func buySEgg(ctx context.Context) {
	fmt.Println("去买SEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买SEgg了")
		return
	default:

	}
	fmt.Println("买SE蛋")
}
func buyBEgg(ctx context.Context) {
	fmt.Println("去买BEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买BEgg了")
		return
	default:

	}
	fmt.Println("买BE蛋")
}
