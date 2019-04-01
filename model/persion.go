package model

import (
	"fmt"
)

type Persion struct {
	//gorm.Model
	FirstName string `json:first_name`
	LastName  string `json:last_name`
}

func (p *Persion) TableName() string {
	fmt.Print("开始被执行！")
	return "person"
}
