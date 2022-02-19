package main

import "fmt"

func main() {
	var a Open = Refrigerator{}
	var b Close = Refrigerator{}
	fmt.Println(a.Open())
	fmt.Println(b.Close())

	var c Box = Refrigerator{}
	fmt.Println(c.Open())

}

type Box interface {
	Open
	Close
}

type Open interface {
	Open() string
}
type Close interface {
	Close() string
}

type Refrigerator struct {
	Size string
}

func (Refrigerator) Open() string {
	return "开"
}
func (Refrigerator) Close() string {
	return "关"
}
