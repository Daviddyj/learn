package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}

	for i := 0; i < 100; i++ {
		go func() {
			m.Store(i, 0)
			for {
				v, ok := m.Load(i)
				if !ok {
					continue
				}
				m.Store(i, v.(int)+1)

				//v := m[i]
				//m[i] = v
				fmt.Println("i=", v)
			}
		}()
	}
	time.Sleep(1 * time.Millisecond)
}
