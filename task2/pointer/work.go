package main

import "fmt"

func main() {
	num := 123
	changeNum(&num)
	fmt.Println(num) // 133

	sli := []int{1, 2, 3, 4, 5}
	changeSlice(&sli)
	fmt.Println(sli) // [2 4 6 8 10]
}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func changeNum(ptr *int) {
	*ptr += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func changeSlice(ptr *[]int) {
	for i := range *ptr {
		(*ptr)[i] *= 2
	}
}
