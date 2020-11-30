package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/controller"
)

func TokenRoutes(r *gin.RouterGroup) {
	r.POST("/refresh", controller.RefreshToken)
}
