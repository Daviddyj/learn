package calclator

import "errors"

func CalcBmi(weight float64, tall float64) (bmi float64, err error) {
	if tall <= 0 {
		//panic("身高不能是0或者负数")
		//err =
		return 0, errors.New("身高不能为负数")
	}
	//todo 验证体重的合法性
	bmi = weight / (tall * tall)
	return bmi, nil
}
