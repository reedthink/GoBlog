package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/reedthink/pkg/setting"
	"github.com/reedthink/routers/api/v1"

)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTag)
		apiv1.PUT("/tags/:id",v1.EditTag)
		apiv1.DELETE("/tags/:id",v1.DeleteTag)

	}

	return r
}
