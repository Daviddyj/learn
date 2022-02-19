package calclator

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBmi(weight float64, tall float64) (bmi float64) {
	bmi, _ = gobmi.BMI(weight, tall)
	return bmi

}
