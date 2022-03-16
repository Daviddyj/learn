package main

import (
	"learn/pkg/apis"
	"math"
	"sort"
)

var _ ServerInterface = &FatRateRank{}

type RandItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	Item []RandItem
}

func (f *FatRateRank) RegisterPersonInformation(pi *apis.PersonInformation) error {
	//TODO implement me
	panic("implement me")
}

func (f *FatRateRank) UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FatRateRank) GetFateRate(name string) (*apis.PersonRank, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FatRateRank) GetTop() ([]*apis.PersonRank, error) {
	//TODO implement me
	panic("implement me")
}

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{
		Item: make([]RandItem, 0, 100),
	}
}

var personFatRate = map[string]float64{}

func (f *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, value := range fatRate {
		if value <= minFatRate {
			minFatRate = value
		}
	}
	for i, item := range f.Item {
		if item.Name == name {
			item.FatRate = minFatRate
		}
		f.Item[i] = item //?  啥意思
	}
	personFatRate[name] = minFatRate
}

func getRand(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fataRate = frItem
				return
			}

		}

	}
	return 0, 0
}
