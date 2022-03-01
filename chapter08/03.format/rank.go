package main

import (
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}
type FatRateRank struct {
	Item []RankItem
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

func (f *FatRateRank) getRand(name string) (rank int, fataRate float64) {
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
