package main

import (
	"fmt"
	"sync"
)

//栈

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

type Stack struct {
	wg   sync.Mutex
	data []interface{}
}

func (q *Stack) Push(data interface{}) {
	q.wg.Lock()
	defer q.wg.Unlock()
	q.data = append([]interface{}{data}, q.data...)
}

func (q *Stack) Pop() (data interface{}, bool bool) {
	q.wg.Lock()
	defer q.wg.Unlock()
	if len(q.data) > 0 {
		num := q.data[0]
		q.data = q.data[1:]
		return num, true
	}
	return nil, false
}
