package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/l6l6ng/go-web2/utils"
	"net/http"
)

func InitRoute() {
	gin.SetMode(utils.AppMode)

	routes := gin.Default()

	routes.Group("api/v1")
	{
		routes.GET("hello", func(context *gin.Context) {
			context.AsciiJSON(http.StatusOK, gin.H{
				"msg": "hello",
			})
		})
	}
	routes.Run(utils.HttpPort)
}
