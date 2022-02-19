package calclator

import go_bmi "learn.go/staging/src/github.com/armstrongli/go-bmi"

func CalcFatRate(age int, bmi float64, sex string) (fatRate float64) {
	return go_bmi.CalcFatRate(age, bmi, sex)

}

//新增需求   根据年龄不同   他的年龄权重也不同
func ageWeight(age int) (per float64) {
	per = 0.23
	if age > 30 && age <= 40 {
		per = 0.22
	}
	return
}
