package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
)

func LoginRoutes(r *gin.RouterGroup) {
	r.POST("/", controller.Login)
}
