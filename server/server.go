package server

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/config"
	"github.com/meastblue/godo/route"
)

func Run() {
	config.Init("yaml", "env", "./config")
	g := gin.Default()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	route.Init(g)

	g.Run()
}
