package apis

//type PersonInformation struct {
//	Name   string  `json:"name"`
//	Sex    string  `json:"sex"`
//	Tall   float64 `json:"tall"`
//	Weight float64 `json:"weight"`
//	Age    int     `json:"age"`
//}
<<<<<<< HEAD
type PersonInformationFatRate struct {
	Name    string
	FatRate float64
}

type PersonRank struct {
	Name       string
	RankNumber int
	FatRate    float64
=======

//type PersonInformationFatRate struct {
//	Name    string
//	FatRate float64
//}
//
//type PersonRank struct {
//	Name       string
//	Sex        string
//	RankNumber int
//	FatRate    float64
//}

func (*PersonInformation) TableName() string {
	return "personal_information"
>>>>>>> 0efc8eb... 0327
}
