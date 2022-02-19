package main

type fakeInput struct {
}

func (*fakeInput) GetInput() Person {

	return Person{
		name:   "daiyijie",
		sex:    "男",
		tall:   1.75,
		weight: 75,
		age:    24,
	}
}
