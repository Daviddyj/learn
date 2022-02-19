package main

import "fmt"

func main() {
	security := Assets{assets: []Asset{&WoodDoor{}}}
	fmt.Println("开始上班")
	security.DoStartWork()
	fmt.Println("8小时，准备下班")
	security.DoStopWork()

}
