package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:pw@tcp(xx:3307)/blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	//defer DB.Close()


	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(20)

	if err := DB.DB().Ping(); err != nil{
		log.Fatalln(err)
	}
}

