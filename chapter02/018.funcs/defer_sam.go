package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time:", startTime)
	defer fmt.Println("duration:", time.Now().Sub(startTime)) //时间为4.72ms   defer准备的是fmt
	defer func() {
		fmt.Println("duration--update:", time.Now().Sub(startTime)) //时间为5.00472s    defer准备的是func()函数
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("finish time:", time.Now())
}

func openFile() {
	fileName := "C:\\Users\\Jack\\Desktop\\dyj.txt"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(0)
	}
	defer fileObj.Close()
}
