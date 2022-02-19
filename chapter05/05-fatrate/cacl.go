package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct {
}

func (Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return err
	}
	person.bmi = bmi
	return nil
}
func (Calc) FatRate(person *Person) error {
	fatRate := gobmi.CalcFatRate(person.age, person.bmi, person.sex)
	person.fatRate = fatRate
	return nil
}
