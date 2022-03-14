package main

import (
	"fmt"
	"learn/pkg/apis"
	"math/rand"
	"time"
)

type Interface interface {
	ReadPersonInformation() (apis.PersonInformation, error)
}

type fakePersonInformation struct {
}

type randPersonInformation struct {
}

func (f *fakePersonInformation) ReadPersonInformation() (apis.PersonInformation, error) {
	return apis.PersonInformation{
		Name:   "戴一杰",
		Sex:    "男",
		Tall:   float32(1.75),
		Weight: float32(76),
		Age:    int64(26),
	}, nil
}

func (f *randPersonInformation) ReadPersonInformation() (apis.PersonInformation, error) {
	rand.Seed(time.Now().UnixNano())
	return apis.PersonInformation{
		Name:   fmt.Sprintf("%f", randFloats(0.01, 0.9, 1)),
		Sex:    "男",
		Tall:   float32(1.75) + rand.Float32(),
		Weight: float32(76) + randFloats(1.0, 5.6, 1),
		Age:    int64(26),
	}, nil
}
func randFloats(min, max float32, n int) float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = min + rand.Float32()*(max-min)
	}
	result := res[0]
	return result
}
