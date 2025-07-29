package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:wcq5201314@tcp(127.0.0.1:3306)/wcq?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("表创建失败: " + err.Error())
	}
	// 编写Go代码，使用Gorm查询某个用户(假设是用户1)发布的所有文章及其对应的评论信息。
	userID := 1
	posts := []Post{}
	db.Where("user_id = ?", userID).
		Preload("Comments"). // 加载关联评论
		Find(&posts)
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	post := Post{}
	db.Order("comment_count DESC").First(&post)

}

type User struct {
	gorm.Model
	// 文章数量
	PostCount int `gorm:"default:0"`
	Post      []Post
}

type Post struct {
	gorm.Model
	// 评论数量
	CommentCount int `gorm:"default:0"`
	// 评论状态
	CommentStatus string `gorm:"default:'无评论'"`
	UserID        uint
	Comment       []Comment
}

type Comment struct {
	gorm.Model
	PostID uint
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 原子操作更新用户文章计数
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1"))

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	// 查询关联文章
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}

	// 统计剩余评论数量
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	// 更新文章状态
	status := "有评论"
	if count == 0 {
		status = "无评论"
	}

	// 原子操作更新文章状态
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  count,
			"comment_status": status,
		}).Error; err != nil {
		return err
	}
	return nil
}
