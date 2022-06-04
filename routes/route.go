package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/l6l6ng/go-web2/api/v1"
	"github.com/l6l6ng/go-web2/utils"
)

func InitRoute() {
	gin.SetMode(utils.AppMode)

	routes := gin.Default()

	routers := routes.Group("api/v1")
	{
		//用户模块
		routers.POST("user/add", v1.AddUser)
		routers.GET("users", v1.GetUsers)
		routers.PUT("user/:id", v1.EditUser)
		routers.DELETE("user/:id", v1.DeleteUser)

		//分类模块

		//文章模块
	}
	routes.Run(utils.HttpPort)
}
