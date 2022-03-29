package main

import (
	"fmt"
	"gorm.io/gorm"
	"learn/chapter12/02.pratice/frinterface"
	"learn/pkg/apis"
	"log"
)

var _ frinterface.ServerInterface = &dbRank{}
var _ frinterface.RankInitInterface = &dbRank{}

func NewDbRank(conn *gorm.DB, embedRank frinterface.ServerInterface) frinterface.ServerInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	if embedRank == nil {
		log.Fatal("内存排行榜为空")
	}
	return &dbRank{
		conn:      conn,
		embedRank: embedRank,
	}
}

type dbRank struct {
	conn      *gorm.DB
	embedRank frinterface.ServerInterface
}

func (d dbRank) Init() error {
	output := make([]*apis.PersonInformation, 0)
	resp := d.conn.Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return resp.Error
	}
	for _, item := range output {
		if _, err := d.embedRank.UpdatePersonInformation(item); err != nil {
			log.Fatalf("初始化%s时失败:%v", item.Name, err)
		}
	}
	return nil
}

func (d dbRank) RegisterPersonInformation(pi *apis.PersonInformation) error {
	resp := d.conn.Create(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("创建%s时失败%v:", pi.Name, err)
		return err
	}
	fmt.Printf("创建%s成功\n", pi.Name)
	_ = d.embedRank.RegisterPersonInformation(pi) //handle error
	return nil
}

func (d dbRank) UpdatePersonInformation(pi *apis.PersonInformation) (*apis.PersonInformationFatRate, error) {
	resp := d.conn.Updates(pi)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%s时失败%v:", pi.Name, err)
		return nil, err
	}
	fmt.Printf("更新%s成功\n", pi.Name)
	return d.embedRank.UpdatePersonInformation(pi)

}

func (d dbRank) GetFateRate(name string) (*apis.PersonRank, error) {
	return d.embedRank.GetFateRate(name)
}

func (d dbRank) GetTop() ([]*apis.PersonRank, error) {
	return d.embedRank.GetTop()
}
