package main

type PersonInformation struct {
	Id     int     `json:"id,omitempty" gorm:"column:id"`
	Name   string  `json:"name,omitempty" gorm:"column:name"`
	Sex    string  `json:"sex,omitempty" gorm:"column:sex"`
	Tall   float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age    int64   `json:"age,omitempty" gorm:"column:age"`
}

func (*PersonInformation) TableName() string {
	return "personal_information"
}
