package main

import "fmt"

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
func main() {
	emp := Employee{Person{"张三", 28}, 1001}
	emp.PrintInfo()
	/**
	员工ID: 1001
	姓名: 张三
	年龄: 28
	*/
}

// person结构体
type Person struct {
	Name string
	Age  int
}

// employee结构体
type Employee struct {
	Person
	EmployeeID int
}

func (emp *Employee) PrintInfo() {
	fmt.Printf("员工ID: %d\n姓名: %s\n年龄: %d\n",
		emp.EmployeeID, emp.Name, emp.Age)
}
