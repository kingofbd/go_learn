package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "缺少认证token")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "无效的token")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
