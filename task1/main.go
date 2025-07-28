package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 5, 1, 5, 4, 4, 3}
	result := f1(arr)
	fmt.Println(result)

	bool1 := f2(12327)
	fmt.Println(bool1)

	bool2 := f3("{[](})")
	fmt.Println(bool2)

	strs := []string{"flower", "flow", "flight"}
	prefix := f4(strs)
	fmt.Println(prefix)

	arr2 := []int{9, 9, 9}
	result2 := f5(arr2)
	fmt.Println(result2)

	arr3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	num := f6(arr3)
	fmt.Println(num)

	arr4 := [][]int{{1, 3}, {2, 6}, {8, 10}, {10, 18}, {24, 35}}
	result3 := f7(arr4)
	fmt.Println(result3)

	arr5 := []int{2, 7, 11, 15}
	result4 := f8(arr5, 9)
	fmt.Println(result4)
}

// 只出现一次的数字
func f1(arr []int) int {
	result := 0
	for _, v := range arr {
		// 通过异或操作抵消成对的数字，比如2^5^2=5
		result ^= v
	}
	return result
}

// 回文数
func f2(x int) bool {
	// 特殊情况处理。1、负数不是回文。2、个位数为0的也不可能是回文
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	half := 0
	// 反转数字的后半部分，直到原始数字小于等于反转部分
	for x > half {
		half = half*10 + x%10
		x /= 10
	}
	// 偶数长度，反转部分=剩余部分，比如1221（12=12）
	// 奇数长度，剩余部分=反转部分/10，比如12321（12==123/10）
	return x == half || x == half/10
}

// 有效的括号
func f3(s string) bool {
	// 使用rune切片作为栈
	stack := make([]rune, 0, len(s))
	// 建立括号映射关系：右括号->左括号
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, char)
		case ')', '}', ']':
			// 检查栈状态和匹配情况
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			// 匹配成功则弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 检查是否所有括号都已匹配
	return len(stack) == 0
}

// 最长公共前缀
func f4(strs []string) string {
	// 处理边界情况
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// 初始化公共前缀为第一个字符串
	prefix := strs[0]

	// 横向扫描：依次与后续字符串比较
	for i := 1; i < len(strs); i++ {
		// 当前字符串与当前前缀的匹配位置
		j := 0
		minLen := min(len(prefix), len(strs[i]))

		// 逐个字符比较直到不匹配
		for j < minLen && prefix[j] == strs[i][j] {
			j++
		}

		// 更新公共前缀
		prefix = prefix[:j]

		// 提前终止：前缀已为空
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// 辅助函数：获取最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 加一
func f5(digits []int) []int {
	// 从最低位开始向高位遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位小于9，直接加1并返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 当前位是9，进位后置0
		digits[i] = 0
	}
	// 处理全为9的特殊情况：在数组前添加1
	return append([]int{1}, digits...)
}

// 删除有序数组中的重复项
func f6(nums []int) int {
	// 边界条件：空数组或单元素数组直接返回
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0 // 慢指针（记录唯一元素位置）
	// 快指针遍历数组（从第二个元素开始）
	for fast := 1; fast < len(nums); fast++ {
		// 发现新元素时，慢指针前移并更新值
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	// 唯一元素数量 = 慢指针位置 + 1
	return slow + 1
}

// 合并区间
func f7(intervals [][]int) [][]int {
	// 处理边界情况
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 按区间起始值排序（核心预处理）
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果集并添加首个区间
	merged := [][]int{intervals[0]}

	// 3. 遍历处理后续区间
	for _, current := range intervals[1:] {
		last := merged[len(merged)-1] // 获取结果集最后一个区间

		// 4. 检查重叠：当前区间起始值 ≤ 上一区间结束值
		if current[0] <= last[1] {
			// 5. 合并区间：扩展结束值到两者最大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 6. 无重叠：添加新区间到结果集
			merged = append(merged, current)
		}
	}

	return merged
}

// 两数之和
func f8(nums []int, target int) []int {
	// 创建哈希表：key=数值, value=索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算互补值
		complement := target - num

		// 检查互补值是否在哈希表中
		if j, ok := numMap[complement]; ok {
			return []int{j, i} // 找到解，返回索引对
		}

		// 存入当前值及其索引
		numMap[num] = i
	}

	return nil // 无解返回 nil
}
