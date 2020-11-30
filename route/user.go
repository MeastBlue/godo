package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
	"github.com/meastblue/godo/middleware"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", middleware.TokenAuthMiddleware(), controller.GetUsers)
	r.GET("/:id", middleware.TokenAuthMiddleware(), controller.GetUser)
	r.POST("/", controller.AddUser)
	r.PUT("/", middleware.TokenAuthMiddleware(), controller.UpdateUser)
	r.DELETE("/:id", middleware.TokenAuthMiddleware(), controller.DeleteUser)
}
