package handler

import (
	"fmt"
	"gin-example/helper"
	"gin-example/model"
	"github.com/gin-gonic/gin"
)

func PutData(c *gin.Context) {
	var persion model.Persion
	fmt.Print(c.Query("firstname"))
	persion.FirstName = c.Query("firstname")
	persion.LastName = c.Query("lastname")
	err := model.PutOne(&persion)

	if err != nil {
		helper.RespondJSON(c, 404, persion)
	} else {
		helper.RespondJSON(c, 200, persion)
	}

}

func GetData(c *gin.Context) {
	var persion []model.Persion
	err := model.GetAll(&persion)
	if err != nil {
		helper.RespondJSON(c, 404, persion)
	} else {
		helper.RespondJSON(c, 200, persion)
	}

}
