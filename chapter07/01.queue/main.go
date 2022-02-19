package main

import (
	"fmt"
	"sync"
)

//实现一个线程安全的队列

func main() {
	q := Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

type Queue struct {
	wg   sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data interface{}) {
	q.wg.Lock()
	defer q.wg.Unlock()
	q.data = append(q.data, data)

}

func (q *Queue) Pop() (data interface{}, bool bool) {
	q.wg.Lock()
	defer q.wg.Unlock()
	if len(q.data) > 0 {
		num := q.data[0]
		q.data = q.data[1:]
		return num, true
	}
	return nil, false
}
