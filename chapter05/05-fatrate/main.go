package main

func main() {

	frSvc := &fatRateService{S: GetFatRateSuggestion(),
		input:  &fakeInput{},
		output: &StdOut{}}

	for {
		p := frSvc.input.GetInput()
		frSvc.GiveSuggestionToPerson(&p)
	}

}
