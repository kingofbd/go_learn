package controllers

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/database"
	"go_learn/task4/models"
	"go_learn/task4/utils"
	"strconv"
)

func CreateComment(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, 400, "无效的文章ID格式")
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.ErrorResponse(c, 400, "无效的请求参数")
		return
	}

	comment.UserID = userID
	comment.PostID = uint(postID) // 类型转换为uint

	if err := database.DB.Create(&comment).Error; err != nil {
		utils.ErrorResponse(c, 500, "评论创建失败")
		return
	}

	utils.SuccessResponse(c, comment)
}

func GetComments(c *gin.Context) {
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, 400, "无效的文章ID格式")
		return
	}

	var comments []models.Comment
	if err := database.DB.Preload("User").
		Where("post_id = ?", uint(postID)). // 类型转换为uint
		Find(&comments).Error; err != nil {
		utils.ErrorResponse(c, 500, "评论获取失败")
		return
	}
	utils.SuccessResponse(c, comments)
}
