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

	// 创建员工表
	sql_create := "create table if not exists employees (\n  id int unsigned auto_increment primary key,\n  name varchar(50) not null comment '员工姓名',\n  department varchar(50) not null comment '所属部门',\n  salary decimal(10,2) not null comment '薪资',\n  created_at timestamp default current_timestamp\n) engine=innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci;"
	db.Exec(sql_create)
	// 向员工表中插入几条数据
	emps := []Employee{
		{Name: "张三", Department: "技术部", Salary: 15000.50},
		{Name: "李四", Department: "市场部", Salary: 12000.75},
		{Name: "王五", Department: "技术部", Salary: 18000.25},
		{Name: "赵六", Department: "人力资源", Salary: 11000.00},
	}
	sql_insert := "insert into employees (name,department,salary) values (?,?,?)"
	stmt, err := db.Preparex(sql_insert)
	if err != nil {
		panic(err)
	}
	for _, emp := range emps {
		_, err := stmt.Exec(emp.Name, emp.Department, emp.Salary)
		if err != nil {
			panic(err)
		}
	}

	res := f1(db)
	fmt.Println(res)
	res2 := f2(db)
	fmt.Println(res2)

	defer db.Close()
}

// 员工信息
type Employee struct {
	ID         uint
	Name       string
	Department string
	Salary     float64
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func f1(db *sqlx.DB) []Employee {
	sqlstr := "select name,department,salary from employees where department = ?"
	res := []Employee{}
	err := db.Select(&res, sqlstr, "技术部")
	if err != nil {
		panic(err)
	}
	return res
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func f2(db *sqlx.DB) Employee {
	sqlstr := "select name,department,salary from employees order by salary desc limit 1"
	res := Employee{}
	err := db.Get(&res, sqlstr)
	if err != nil {
		panic(err)
	}
	return res
}
