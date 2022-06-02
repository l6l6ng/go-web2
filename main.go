package main

import "github.com/l6l6ng/go-web2/routes"
import "github.com/l6l6ng/go-web2/model"

func main() {
	model.InitDb()
	routes.InitRoute()
}
