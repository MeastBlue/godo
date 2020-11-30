package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
	"github.com/meastblue/godo/middleware"
)

func TaskRoutes(r *gin.RouterGroup) {
	r.GET("/", middleware.TokenAuthMiddleware(), controller.GetTasks)
	r.GET("/:id", middleware.TokenAuthMiddleware(), controller.GetTask)
	r.POST("/", middleware.TokenAuthMiddleware(), controller.AddTask)
	r.PUT("/:id", middleware.TokenAuthMiddleware(), controller.UpdateTask)
	r.DELETE("/:id", middleware.TokenAuthMiddleware(), controller.DeleteTask)
}
