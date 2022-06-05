package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l6l6ng/go-web2/routes"
	"io"
	"os"
)
import "github.com/l6l6ng/go-web2/model"

func main() {
	model.InitDb()
	routes.InitRoute()

	//记录日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
