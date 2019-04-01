package model

import (
	"fmt"
	"gin-example/config"
)

func GetAll(p *[]Persion) error {

	err := config.DB.Find(p).Error
	if err != nil {
		return err
	}

	return nil
}

func PutOne(p *Persion) (err error) {

	fmt.Println(p)
	err = config.DB.Save(p).Error
	if err != nil {
		return err
	}
	return nil
}
