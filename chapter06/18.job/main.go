package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Second)
	defer cancel()

	flag := make(chan bool, 1)
	go account(ctx)
	go configure(ctx)
	go distributeService(ctx)
	go verifyService(ctx, flag)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时，没有完成")
	default:
		for v := range flag {
			if v {
				fmt.Println("所有任务最终部署完成")
			} else {
				fmt.Println("所有任务部署失败，任务下线或者任务进行重试")
			}
			close(flag)
		}
	}
	//time.Sleep(7 * time.Second) //这边是不是不要要sleep     因为select选择器本身就会等待  难怪刚刚7秒钟过去了程序还没结束
}

func account(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	doneCh := make(chan string, 2)
	time.Sleep(5 * time.Second)
	go accountRegister(ctx, doneCh)
	go accountGrantAuth(ctx, doneCh)

	successAccount := 0
	for v := range doneCh {
		successAccount++
		fmt.Println("job:", v, "Done")
		if successAccount == 2 {
			close(doneCh)
		}
	}

	fmt.Println("账号处理完成")

}
func accountRegister(ctx context.Context, doneCh chan string) {
	fmt.Println("注册账号")
	defer fmt.Println("注册账号完成")

	for { //类似与拨打电话，一般重试两到三次，确保服务的成功部署
		//调用...接口
		select {
		case <-ctx.Done():
			fmt.Println("context结束,取消注册账号")
			return
		default:
		}
		doneCh <- "accountRegister 完成"
		break
	}

}
func accountGrantAuth(ctx context.Context, doneCh chan string) {
	fmt.Println("授权账号")
	defer fmt.Println("授权账号完成")
	for { //类似与拨打电话，一般重试两到三次，确保服务的成功部署
		//调用...接口
		select {
		case <-ctx.Done():
			fmt.Println("context结束,取消授权账号")
			return
		default:
		}
		doneCh <- "accountGrantAuth 完成"
		break
	}

}

func distributeService(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 7*time.Minute)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go distributeInstance(ctx, &wg)
	go distributeLB(ctx, &wg)
	wg.Wait()
	fmt.Println("distributeService Done")

}
func distributeInstance(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束,取消部署实例")
			return
		default:
		}
		fmt.Println("部署实例")
		break
	}

}
func distributeLB(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束,取消部署负载均衡器")
			return
		default:
		}
		fmt.Println("部署负载觉衡器")
		break
	}
}

func configure(_ context.Context) {
	fmt.Println("新账号注入")
}

func verifyService(ctx context.Context, flag chan bool) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go verifyFunction(ctx, &wg)
	wg.Wait()
	fmt.Println("verifyService Done")
	//time.Sleep(15 * time.Second)
	flag <- true
}

func verifyFunction(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("取消验证服务")
		default:
		}
		fmt.Println("开始验证服务")
		time.Sleep(1000 * time.Millisecond) //用来替换服务调用，服务验证
		if i <= 1 {
			fmt.Println("再一次进行验证")
			continue
		}
		fmt.Println("验证完成")
		break
	}

}
