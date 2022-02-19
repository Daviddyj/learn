package gobmi

func CalcFatRate(age int, bmi float64, sex string) (fatRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + ageWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

//新增需求   根据年龄不同   他的年龄权重也不同
func ageWeight(age int) (per float64) {
	per = 0.23
	if age > 30 && age <= 40 {
		per = 0.22
	}
	return
}
