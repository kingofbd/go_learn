package routes

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/controllers"
	"go_learn/task4/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	// 无需认证的路由
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	// 需要认证的路由
	apiGroup := router.Group("/api")
	apiGroup.Use(middlewares.AuthMiddleware())
	{
		// 文章路由
		postsGroup := apiGroup.Group("/posts")
		{
			postsGroup.POST("", controllers.CreatePost)
			postsGroup.GET("", controllers.GetPosts)
			postsGroup.GET("/:id", controllers.GetPost)
			postsGroup.PUT("/:id", controllers.UpdatePost)
			postsGroup.DELETE("/:id", controllers.DeletePost)

			// 评论路由
			postsGroup.POST("/:postId/comments", controllers.CreateComment)
			postsGroup.GET("/:postId/comments", controllers.GetComments)
		}
	}
}
