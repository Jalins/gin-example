package routers

import (
	"gin-example/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.POST("/put", handler.PutData)
		v1.GET("/get", handler.GetData)
	}

	return r
}
