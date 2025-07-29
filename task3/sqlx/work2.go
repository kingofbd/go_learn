package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// 连接mysql
	dsn := "root:wcq5201314@tcp(127.0.0.1:3306)/wcq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 创建书表
	sql_create := "create table if not exists books (\n  id int unsigned auto_increment primary key,\n  title varchar(50) not null comment '书名',\n  author varchar(50) not null comment '作者',\n  price decimal(10,2) not null comment '定价',\n  created_at timestamp default current_timestamp\n) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;"
	db.Exec(sql_create)
	// 向书表中插入几条数据
	books := []Book{
		{Title: "世界简史", Author: "张三", Price: 150.50},
		{Title: "中国简史", Author: "李四", Price: 188.50},
		{Title: "时间简史", Author: "霍金", Price: 254.50},
	}
	sql_insert := "insert into books (title, author, price) values (?, ?, ?)"
	stmt, err := db.Preparex(sql_insert)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		_, err := stmt.Exec(book.Title, book.Author, book.Price)
		if err != nil {
			panic(err)
		}
	}
	res := f3(db)
	fmt.Println(res)
}

// 定义一个 Book 结构体，包含与 books 表对应的字段。
type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func f3(db *sqlx.DB) []Book {
	str := "select title,author,price from books where price > ?"
	res := []Book{}
	err := db.Select(&res, str, 50)
	if err != nil {
		panic(err)
	}
	return res
}
