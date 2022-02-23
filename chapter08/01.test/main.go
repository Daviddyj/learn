package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//rand.Seed(time.Now().Unix())
	//var Reader io.Reader
	var a = rand.Reader
	b, _ := rand.Int(a, big.NewInt(55))
	fmt.Println(b)
}
