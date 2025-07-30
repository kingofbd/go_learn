package controllers

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/database"
	"go_learn/task4/models"
	"go_learn/task4/utils"
)

func CreatePost(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.ErrorResponse(c, 400, "无效的请求参数")
		return
	}

	post.UserID = userID
	if err := database.DB.Create(&post).Error; err != nil {
		utils.ErrorResponse(c, 500, "文章创建失败")
		return
	}

	utils.SuccessResponse(c, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Preload("User").Find(&posts).Error; err != nil {
		utils.ErrorResponse(c, 500, "文章获取失败")
		return
	}
	utils.SuccessResponse(c, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Preload("User").Preload("Comments.User").First(&post, id).Error; err != nil {
		utils.ErrorResponse(c, 404, "文章不存在")
		return
	}
	utils.SuccessResponse(c, post)
}

func UpdatePost(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.ErrorResponse(c, 404, "文章不存在")
		return
	}

	if post.UserID != userID {
		utils.ErrorResponse(c, 403, "无权修改此文章")
		return
	}

	var updateData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, 400, "无效的请求参数")
		return
	}

	if updateData.Title != "" {
		post.Title = updateData.Title
	}
	if updateData.Content != "" {
		post.Content = updateData.Content
	}

	if err := database.DB.Save(&post).Error; err != nil {
		utils.ErrorResponse(c, 500, "文章更新失败")
		return
	}

	utils.SuccessResponse(c, post)
}

func DeletePost(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.ErrorResponse(c, 404, "文章不存在")
		return
	}

	if post.UserID != userID {
		utils.ErrorResponse(c, 403, "无权删除此文章")
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		utils.ErrorResponse(c, 500, "文章删除失败")
		return
	}

	utils.SuccessResponse(c, nil)
}
