package frinterface

import "learn/pkg/apis"

type ServerInterface interface {
	RegisterPersonInformation(pi *apis.PersonInformation) error
	UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error)
	GetFateRate(name string) (*apis.PersonRank, error)
	GetTop() ([]*apis.PersonRank, error)
}
<<<<<<< HEAD
=======
type RankInitInterface interface {
	Init() error
}
>>>>>>> 0efc8eb... 0327
