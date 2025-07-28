package main

import (
	"fmt"
	"math"
)

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
func main() {
	// 矩形
	rect := Rectangle{10, 15}
	rect_area := rect.Area()
	rect_perimeter := rect.Perimeter()
	fmt.Println(rect_area, rect_perimeter)
	// 圆形
	circle := Circle{10}
	circle_area := circle.Area()
	circle_perimeter := circle.Perimeter()
	fmt.Println(circle_area, circle_perimeter)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 矩形结构体
type Rectangle struct {
	Width, Height float64
}

// 矩形结构体实现周长和面积的接口
func (rect *Rectangle) Area() float64 {
	return rect.Width * rect.Height
}
func (rect *Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

// 圆形结构体
type Circle struct {
	Radius float64
}

// 圆形结构体实现周长和面积的接口
func (circle *Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}
