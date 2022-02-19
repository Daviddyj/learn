package main

import "fmt"

type shipLengend struct {
}

func (*shipLengend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用ship做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*shipLengend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用ship做 PutElephantIntoRefrigerator")
	return nil
}
func (*shipLengend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用ship做 CloseTheDoorOfRefrigerator")
	return nil
}
