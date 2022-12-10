package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func main() {

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}else {
		fmt.Println("DB接続完了")
	}
	defer db.Close()

	error := db.Create(&Posts{
			author:     "chacha",
			post:      1,
			name:  "lesson1",
			intro: "introintro",
			body: "bodybodt",
		}).Error
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println("データ追加成功")
	}
}


//データベース接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "chacha0503"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "goProject"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	return gorm.Open(DBMS, CONNECT)
}

type Posts struct {
	author string
	post int
	name string
	intro string
	body string
}