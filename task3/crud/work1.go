package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:wcq5201314@tcp(127.0.0.1:3306)/wcq?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// 创建学生表
	stu := Student{}
	db.AutoMigrate(stu)

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。（可以插入多条）
	stus := []Student{
		{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 19, Grade: "二年级"},
		{Name: "王五", Age: 13, Grade: "一年级"},
	}
	db.CreateInBatches(stus, 10)

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	res := []Student{}
	db.Debug().Where("age > ?", 18).Find(&res)
	fmt.Println(res)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Debug().Where("age < ?", 15).Unscoped().Delete(&Student{})
}

// 学生
type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}
