package middlewares

import (
	"example/connectify/app/constant"
	"example/connectify/app/pkg"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer pkg.PanicHandler(c)
		err := pkg.TokenValid(c)
		if err != nil {
			pkg.PanicException(constant.Unauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
