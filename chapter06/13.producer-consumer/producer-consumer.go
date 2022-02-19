package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 10 {
		fmt.Println("生产者看到库存满了，不生产")
		s.pCond.Wait()
	}
	fmt.Println("开始生产,库存为:", s.DataCount)
	s.DataCount++
	s.cCond.Signal()
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到库存为零，不消费")
		s.cCond.Wait()
	}
	fmt.Println("开始消费,库存为:", s.DataCount)
	s.DataCount--
	s.pCond.Signal()
}

func main() {
	Num := &Store{Max: 10}
	pro, con := 50, 50
	Num.pCond = sync.NewCond(&Num.lock) //必须要初始化pCond和cCond
	Num.cCond = sync.NewCond(&Num.lock) //必须要初始化pCond和cCond
	//wg := sync.WaitGroup{}
	//wg.Add(100)
	for i := 0; i < pro; i++ {
		go func(i int) {
			//defer wg.Done()
			for true {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(Num)

			}
		}(i)
	}

	for i := 0; i < con; i++ {
		go func(i int) {
			//defer wg.Done()
			for true {
				time.Sleep(100 * time.Millisecond)

				Consumer{}.Consume(Num)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("最后仓库里还剩下的商品数量为:", Num.DataCount)

}
