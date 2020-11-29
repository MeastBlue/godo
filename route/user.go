package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", controller.GetUsers)
	r.GET("/:id", controller.GetUser)
	r.POST("/", controller.AddUser)
	r.PUT("/", controller.UpdateUser)
	r.DELETE("/:id", controller.DeleteUser)
}
