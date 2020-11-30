package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
	"github.com/meastblue/godo/middleware"
)

func LoginRoutes(r *gin.RouterGroup) {
	r.POST("/", controller.Login)
	r.POST("/logout", middleware.TokenAuthMiddleware(), controller.Logout)
}
