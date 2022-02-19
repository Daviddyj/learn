package main

import "fmt"

type manlengend struct {
}

func (*manlengend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用manLegend做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*manlengend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用manLegend做 PutElephantIntoRefrigerator")
	return nil
}
func (*manlengend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用manLegend做 CloseTheDoorOfRefrigerator")
	return nil
}
