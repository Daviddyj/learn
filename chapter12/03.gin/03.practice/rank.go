package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"learn/chapter12/02.pratice/frinterface"
	"learn/pkg/apis"
	"math"
	"sort"
	"sync"
)

var _ frinterface.ServerInterface = &FatRateRank{}

type RandItem struct {
	Name    string
	Sex     string
	FatRate float64
}

type FatRateRank struct {
	items     []RandItem
	itemsLock sync.Mutex
}

func (f *FatRateRank) RegisterPersonInformation(pi *apis.PersonInformation) error {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		fmt.Println("bmi计算错误:", err)
		return err
	}
	ft := gobmi.CalcFatRate(int(pi.Age), bmi, pi.Sex)
	f.inputRecord(pi.Name, pi.Sex, ft)
	return nil
}

func (f *FatRateRank) UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		fmt.Println("bmi计算错误:", err)
		return nil, err
	}
	ft := gobmi.CalcFatRate(int(pi.Age), bmi, pi.Sex)
	f.inputRecord(pi.Name, pi.Sex, ft)
	//rank, fatRate := getRand(pi.Name)
	return &apis.PersonInformationFatRate{
		Name:    pi.Name,
		FatRate: float32(ft),
	}, err
}

func (f *FatRateRank) GetFateRate(name string) (*apis.PersonRank, error) {
	rank, fatRate := f.getRank(name)
	return &apis.PersonRank{
		Name:       name,
		RankNumber: int64(rank),
		FatRate:    float32(fatRate),
	}, nil
}

func (f *FatRateRank) GetTop() ([]*apis.PersonRank, error) {
	return f.getRankTop(), nil
}

func NewFatRateRank() *FatRateRank {
	return &FatRateRank{
		items: make([]RandItem, 0, 100),
	}
}

//var personFatRate = map[string]float64{}
//
//func (f *FatRateRank) inputRecord(name string, fatRate ...float64) {
//	minFatRate := math.MaxFloat64
//	for _, value := range fatRate {
//		if value <= minFatRate {
//			minFatRate = value
//		}
//	}
//	for i, item := range f.Item {
//		if item.Name == name {
//			item.FatRate = minFatRate
//		}
//		f.Item[i] = item //?  啥意思
//	}
//	personFatRate[name] = minFatRate
//}
//
//func getRand(name string) (rank int, fataRate float64) {
//	fatRate2PersonMap := map[float64][]string{}
//	rankArr := make([]float64, 0, len(personFatRate))
//	for nameItem, frItem := range personFatRate {
//		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
//		rankArr = append(rankArr, frItem)
//	}
//	sort.Float64s(rankArr)
//	for i, frItem := range rankArr {
//		_names := fatRate2PersonMap[frItem]
//		for _, _name := range _names {
//			if _name == name {
//				rank = i + 1
//				fataRate = frItem
//				return
//			}
//
//		}
//
//	}
//	return 0, 0
//}
//
////老师这里用的时切片完成的，而我用的是map   所以最后的浏览器获得的是空，后续可以改进一下，使用切片进行排序，并返回切片
//func (r *FatRateRank) getRankTop() []*apis.PersonRank {
//	r.itemsLock.Lock()
//	defer r.itemsLock.Unlock()
//	sort.Slice(r.Item, func(i, j int) bool {
//		return r.Item[i].FatRate < r.Item[j].FatRate
//	})
//	out := make([]*apis.PersonRank, 0, 10)
//	for i := 0; i < 10 && i < len(r.Item); i++ {
//		item := r.Item[i]
//		out = append(out, &apis.PersonRank{
//			Name:       item.Name,
//			RankNumber: i,
//			FatRate:    item.FatRate,
//		})
//	}
//	return out
func (r *FatRateRank) inputRecord(name, sex string, fatRate ...float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RandItem{
			Name:    name,
			Sex:     sex,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}

	return
}

func (r *FatRateRank) getRankTop() []*apis.PersonRank {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	out := make([]*apis.PersonRank, 0, 10)
	for i := 0; i < 10 && i < len(r.items); i++ {
		item := r.items[i]
		out = append(out, &apis.PersonRank{
			Name:       item.Name,
			Sex:        item.Sex,
			RankNumber: int64(i),
			FatRate:    float32(item.FatRate),
		})
	}
	return out
}
