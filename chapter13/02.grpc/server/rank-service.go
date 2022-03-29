package main

import (
	"context"
	context2 "golang.org/x/net/context"
	"io"
	"learn/chapter13/02.grpc/server/frinterface"
	"learn/pkg/apis"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &rankServer{}

type rankServer struct {
	sync.Mutex
	rankS    frinterface.ServerInterface
	personCh chan *apis.PersonInformation
}

func (r *rankServer) Register(ctx context2.Context, information *apis.PersonInformation) (*apis.PersonInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人:%s\n", information.String())
	return information, nil
}

func (r *rankServer) UpdatePersonInformation(ctx context.Context, information *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	r.regPerson(information)
	return r.rankS.UpdatePersonInformation(information)

}

func (r *rankServer) GetFateRate(ctx context.Context, information *apis.PersonInformation) (*apis.PersonRank, error) {
	return r.rankS.GetFateRate(information.Name)
}

func (r *rankServer) GetTop(ctx context.Context, null *apis.Null) (*apis.PersonRanks, error) {
	top, err := r.rankS.GetTop()
	if err != nil {
		log.Println("获取榜单时错误:", err)
	}
	return &apis.PersonRanks{PersonRanks: top}, nil

}

func (r *rankServer) regPerson(p *apis.PersonInformation) {
	r.rankS.RegisterPersonInformation(p) //todo handle error
	r.personCh <- p
}

func (r *rankServer) WatchPersons(null *apis.Null, server apis.RankService_WatchPersonsServer) error {

	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			if err != nil {
				log.Println("发送失败，结束:", err)
				return err
			}
		}
	}
	return nil
}

func (r *rankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for true {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatal("获取用户输入出错:", err)
		}
		pis.Items = append(pis.Items, pi)
		r.regPerson(pi)
	}
	log.Println("连续得到注册清单:", pis.Items)
	return server.SendAndClose(pis)
}
