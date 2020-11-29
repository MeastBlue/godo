package route

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			UserRoutes(user)
		}
		task := v1.Group("/task")
		{
			TaskRoutes(task)
		}
		login := v1.Group("/login")
		{
			LoginRoutes(login)
		}
	}
}
