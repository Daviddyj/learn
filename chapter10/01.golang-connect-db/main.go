package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn/pkg/apis"
	"log"
)

func main() {
	learnDB, err := getDBConnection()
	defer learnDB.Close()
	err = testDBConnection(err, learnDB)
	//insertData(learnDB)
	queryAllDataWithHack(err, learnDB)

}

func insertData(learnDB *sql.DB) error {
	_, err := learnDB.Exec(fmt.Sprintf("insert into personal_information(id,name,sex,tall,weight,age) values ('%d','%s','%s','%f','%f','%d')",
		888,
		"焦莹",
		"女",
		1.60,
		75.0,
		19))

	if err != nil {
		fmt.Println("插入错误", err)
		return err
	}
	return nil
}

func queryAllData(err error, learnDB *sql.DB) {
	rows, err := learnDB.Query("select name ,age from personal_information ")
	if err != nil {
		log.Fatal(err)
	}
	list := apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Printf("name:%s\tage:%d\n", name, age)
		list.Items = append(list.Items, &apis.PersonInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, err := json.Marshal(list)
	fmt.Println(string(data))
}
func queryAllDataWithHack(err error, learnDB *sql.DB) {
	_sql := fmt.Sprintf("select name ,age from personal_information where name = '%s' and sex ='%s'", "郭杰' -- ", "女")
	rows, err := learnDB.Query(_sql)
	if err != nil {
		log.Fatal(err)
	}
	list := apis.PersonalInformationList{}
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Printf("name:%s\tage:%d\n", name, age)
		list.Items = append(list.Items, &apis.PersonInformation{
			Name: name,
			Age:  int64(age),
		})
	}
	data, err := json.Marshal(list)
	fmt.Println(string(data))
}

func testDBConnection(err error, learnDB *sql.DB) error {
	if err = learnDB.Ping(); nil != err {
		log.Fatal("DB 测试失败:", err)
	}
	fmt.Println("数据库连接成功！可以执行命令了")
	return err
}

func getDBConnection() (*sql.DB, error) {
	learnDB, err := sql.Open("mysql", "root:1165207594dyj@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("链接数据库失败:", err)
	}
	return learnDB, err
}
