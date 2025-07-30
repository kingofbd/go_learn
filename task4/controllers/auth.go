package controllers

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/database"
	"go_learn/task4/models"
	"go_learn/task4/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, 400, "无效的请求参数")
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(c, 500, "密码加密失败")
		return
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, 500, "用户创建失败")
		return
	}

	utils.SuccessResponse(c, gin.H{"id": user.ID})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, "无效的请求参数")
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, 401, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.ErrorResponse(c, 401, "用户名或密码错误")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ErrorResponse(c, 500, "token生成失败")
		return
	}

	utils.SuccessResponse(c, gin.H{"token": token})
}
