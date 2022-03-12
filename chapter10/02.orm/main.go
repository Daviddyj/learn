package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:1165207594dyj@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonInformation{
		//Id:     500,
		Name:   "郭杰123",
		Sex:    "男",
		Tall:   1.65,
		Weight: 70,
		Age:    26,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建人***时失败")
		return err
	}
	fmt.Println("创建人***成功")
	return nil
}
func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonInformation{
		Id:     500, //更新必须要有ID
		Name:   "郭杰123",
		Sex:    "男",
		Tall:   1.65,
		Weight: 70,
		Age:    2000,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败")
		return err
	}
	fmt.Println("更新***成功")
	return nil
}

//部分更新  在生产环境中我们一般都是部分更新
func updateExistingPersonSelectedFields(conn *gorm.DB) error {
	p := &PersonInformation{
		Id:     500, //更新必须要有ID
		Name:   "郭杰戴一杰",
		Sex:    "男",
		Tall:   1.65,
		Weight: 70,
		Age:    2000,
	}
	resp := conn.Model(p).Select("name").Updates(p)

	if err := resp.Error; err != nil {
		fmt.Println("更新人***时失败")
		return err
	}
	fmt.Println("更新***成功")
	return nil
}

func ormSelect(conn *gorm.DB) {
	output := make([]*PersonInformation, 0)
	resp := conn.Where(&PersonInformation{Name: "小王"}).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return
	}
	data, _ := json.Marshal(output)
	fmt.Printf("查询结果:%+v\n", string(data))

	//fmt.Printf("查询结果:%+v\n", output[0])
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) { //模糊查询 并且更加的安全
	output := make([]*PersonInformation, 0)
	resp := conn.Where("age<?", 10).Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return
	}
	data, _ := json.Marshal(output)
	fmt.Printf("查询结果:%+v\n", string(data))
}

//func updateExistingPerson(conn *gorm.DB) error {
//	resp := conn.Create(&PersonInformation{
//		//Id:     500,
//		Name:   "郭杰123",
//		Sex:    "男",
//		Tall:   1.65,
//		Weight: 70,
//		Age:    26,
//	})
//	if err := resp.Error; err != nil {
//		fmt.Println("创建人***时失败")
//		return err
//	}
//	fmt.Println("创建人***成功")
//	return nil
//}

func deletePerson(conn *gorm.DB) error {
	resp := conn.Delete(&PersonInformation{
		Id: 15, //更新必须要有ID
	})
	if err := resp.Error; err != nil {
		fmt.Println("删除人***时失败")
		return err
	}
	fmt.Println("删除***成功")
	return nil
}

func main() {
	conn := connectDb()
	fmt.Println(conn)
	//if err := createNewPerson(conn); err != nil {
	//	//TODO handle err
	//}
	//ormSelectWithInaccurateQuery(conn)
	//if err := updateExistingPerson(conn); err != nil {
	//	//TODO handle err
	//	fmt.Println(err)
	//}

	//if err := updateExistingPersonSelectedFields(conn); err != nil {
	//	//TODO handle err
	//	fmt.Println(err)
	//}
	deletePerson(conn)
}
