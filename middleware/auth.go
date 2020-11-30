package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/util"
)

func TokenAuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
     err := util.TokenValid(c.Request)
     if err != nil {
			util.SendJsonUnauthorized(c, err.Error())
			c.Abort()
        return
     }
     c.Next()
  }
}