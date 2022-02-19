package main

import "log"

type fatRateService struct {
	//c      Calc
	S fatRateSuggestion
	//Person []Person
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
		return "错误"
	}
	person.calFatRate()
	return frsvc.S.GetSuggestion(person)

}

func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range persons {
		out[item] = frsvc.GiveSuggestionToPerson(item)
	}
	return out
}
