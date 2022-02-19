package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Person struct {
	name    string
	sex     string
	tall    float64
	weight  float64
	age     int
	bmi     float64
	fatRate float64
}

func (p *Person) calBmi() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("errpr When calculate for person[%s],%v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calFatRate() {
	fatRate := gobmi.CalcFatRate(p.age, p.bmi, p.sex)
	p.fatRate = fatRate
	fmt.Println(p.fatRate)
}
