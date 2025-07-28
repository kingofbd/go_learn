package main

func main() {

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
