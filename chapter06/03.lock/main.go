package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//for i := 0; i < 10; i++ {
	//countDictLockPrice()
	//
	countDictGoroutineSafeRW1()
	//}
	fmt.Println(time.Now())

}

func forgetUnlockCountDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5)

	for p := 0; p < 5; p++ {
		//fmt.Print("正在统计第", p, "页")
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			//defer totalCountLock.Unlock() //推荐
			totalCount += 100 //注意在，这里有重复覆盖的问题

		}()
	}
	wg.Wait()
	//fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有:", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} //读写锁
	wg := sync.WaitGroup{}
	wg.Add(5)

	for p := 0; p < 5; p++ {
		//fmt.Print("正在统计第", p, "页")
		go func(pageNum int) {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100 //注意在，这里有重复覆盖的问题
			if p == 3 {
				time.Sleep(10 * time.Second)
			}
			totalCountLock.Unlock() //推荐              可以使用defer     但是defer并不能及时的将锁解开   所以用完锁住的内容立马解锁
		}(p)

	}
	wg.Wait()
	//fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有:", totalCount, "字")
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(3)
	for p := 0; p < 3; p++ {
		go func(p int) {
			defer wg.Done()
			totalCountLock.Lock()
			// fmt.Print("正在统计第", p, "页，")

			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			totalCountLock.Unlock()

			time.Sleep(1 * time.Second)

			// totalCountLock.Unlock() // 也可以用defer来解锁
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}
func countDictGoroutineSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	workCount := 5
	wg.Add(workCount)
	doneCh := make(chan struct{})
	for p := 0; p < workCount; p++ {

		go func(p int) {
			for {
				fmt.Println(p, "读锁开始时间", time.Now())
				totalCountLock.RLock()
				fmt.Println(p, "读锁拿到锁时间", time.Now())
				//fmt.Println(totalCount)
				time.Sleep(1 * time.Second)
				totalCountLock.RUnlock()
				//needBreak := false
				//select {
				//case <-doneCh:
				//	needBreak = true
				//default:
				//}
				//if needBreak == true {
				//	break
				//}
			}

		}(p)
	}
	for p := 0; p < workCount; p++ {
		go func(p int) {
			defer wg.Done()
			fmt.Println(p, "写锁开始时间", time.Now())
			totalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间", time.Now())

			// fmt.Print("正在统计第", p, "页，")
			defer totalCountLock.Unlock()

			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题

			// totalCountLock.Unlock() // 也可以用defer来解锁
		}(p)
	}
	wg.Wait()
	close(doneCh)
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}
func countDictGoroutineSafeRW1() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}

	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)

	doneCh := make(chan struct{})
	for p := 0; p < workerCount; p++ {
		go func(p int) { // 读锁可以多个go routine同时拿到。
			fmt.Println(p, "读锁开始时间：", time.Now())
			totalCountLock.RLock()
			fmt.Println(p, "读锁拿到锁时间：", time.Now())
			time.Sleep(1 * time.Second)
			totalCountLock.RUnlock()
		}(p)
	}
	for p := 0; p < workerCount; p++ {
		go func(p int) {
			defer wg.Done()

			fmt.Println(p, "写锁开始时间：", time.Now())
			totalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间：", time.Now())
			// fmt.Print("正在统计第", p, "页，")
			defer totalCountLock.Unlock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			// totalCountLock.Unlock() // 也可以用defer来解锁
		}(p)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}
