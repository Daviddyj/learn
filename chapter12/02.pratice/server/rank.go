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
	FatRate float64
}

type FatRateRank struct {
	Item      []RandItem
	itemsLock sync.Mutex
}

func (f *FatRateRank) RegisterPersonInformation(pi *apis.PersonInformation) error {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		fmt.Println("bmi计算错误:", err)
		return err
	}
	ft := gobmi.CalcFatRate(int(pi.Age), bmi, pi.Sex)
	f.inputRecord(pi.Name, ft)
	return nil
}

func (f *FatRateRank) UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		fmt.Println("bmi计算错误:", err)
		return nil, err
	}
	ft := gobmi.CalcFatRate(int(pi.Age), bmi, pi.Sex)
	f.inputRecord(pi.Name, ft)
	//rank, fatRate := getRand(pi.Name)
	return &apis.PersonInformationFatRate{
		Name:    pi.Name,
<<<<<<< HEAD
		FatRate: ft,
=======
		FatRate: float32(ft),
>>>>>>> 0efc8eb... 0327
	}, err
}

func (f *FatRateRank) GetFateRate(name string) (*apis.PersonRank, error) {
	rank, fatRate := getRand(name)
	return &apis.PersonRank{
		Name:       name,
<<<<<<< HEAD
		RankNumber: rank,
		FatRate:    fatRate,
=======
		RankNumber: int64(rank),
		FatRate:    float32(fatRate),
>>>>>>> 0efc8eb... 0327
	}, nil
}

func (f *FatRateRank) GetTop() ([]*apis.PersonRank, error) {
	return f.getRankTop(), nil
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

//老师这里用的时切片完成的，而我用的是map   所以最后的浏览器获得的是空，后续可以改进一下，使用切片进行排序，并返回切片
func (r *FatRateRank) getRankTop() []*apis.PersonRank {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	sort.Slice(r.Item, func(i, j int) bool {
		return r.Item[i].FatRate < r.Item[j].FatRate
	})
	out := make([]*apis.PersonRank, 0, 10)
	for i := 0; i < 10 && i < len(r.Item); i++ {
		item := r.Item[i]
		out = append(out, &apis.PersonRank{
			Name:       item.Name,
<<<<<<< HEAD
			RankNumber: i,
			FatRate:    item.FatRate,
=======
			RankNumber: int64(i),
			FatRate:    float32(item.FatRate),
>>>>>>> 0efc8eb... 0327
		})
	}
	return out
}
