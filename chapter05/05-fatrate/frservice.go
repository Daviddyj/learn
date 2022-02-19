package main

import "log"

type fatRateService struct {
	//c      Calc
	S      fatRateSuggestion
	input  InputService
	output OutputService
	//Person []Person
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) {
	if err := person.calBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
		return
	}
	person.calFatRate()
	frsvc.output.Output(*person, frsvc.S.GetSuggestion(person))

}

//func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
//	out := map[*Person]string{}
//	for _, item := range persons {
//		out[item] = frsvc.GiveSuggestionToPerson(item)
//	}
//	return out
//}
