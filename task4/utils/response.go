package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	Response(c, code, message, nil)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	Response(c, 200, "success", data)
}
