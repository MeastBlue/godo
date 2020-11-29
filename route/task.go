package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
)

func TaskRoutes(r *gin.RouterGroup) {
	r.GET("/", controller.GetTasks)
	r.GET("/:id", controller.GetTask)
	r.POST("/", controller.AddTask)
	r.PUT("/:id", controller.UpdateTask)
	r.DELETE("/:id", controller.DeleteTask)
}
