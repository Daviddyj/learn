package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Store struct {
	init  sync.Once
	store chan int
	Max   int
}

type Producer struct{}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

func (Producer) Produce(s *Store) {
	s.store <- rand.Int()
	fmt.Println("开始生产+1")
}

type Consumer struct{}

func (Consumer) Consume(s *Store) {
	fmt.Println("消费者消费-1", <-s.store)
}

func main() {
	s := &Store{
		Max: 50,
	}
	s.instrument()
	pCount, cCount := 5, 5
	for i := 0; i < pCount; i++ {
		go func() {
			//for {
			time.Sleep(100 * time.Millisecond)
			Producer{}.Produce(s)
			//}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			//for {
			time.Sleep(100 * time.Millisecond)
			Consumer{}.Consume(s)
			//}
		}()
	}
	time.Sleep(10 * time.Second)
}
