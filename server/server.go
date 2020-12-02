package server

import (
	"fmt"
	"os"

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

	addr := fmt.Sprintf("%s:%s", os.Getenv("srv.Host"), os.Getenv("srv.Port"))
	crt := fmt.Sprintf("%s/%s", os.Getenv("tls.Path"), os.Getenv("tls.Crt"))
	key := fmt.Sprintf("%s/%s", os.Getenv("tls.Path"), os.Getenv("tls.Key"))
	fmt.Println(addr)
	g.RunTLS(addr, crt, key)
}
