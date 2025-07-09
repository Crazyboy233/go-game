package main

import "fmt"

// 1. 创建一个切片，添加一些玩家分数，然后取出前3个分数打印。
func select3() []int {
	a := []int{1, 2, 3, 4, 5}
	return a[:3]
}

// 2. 编写函数 Sum(nums []int) int，返回所有分数之和。
func Sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

// 3. 试试用 copy 和 append 构造一个切片的拷贝。
func mycopy(nums []int) []int {
	// 使用copy函数
	result := make([]int, len(nums))
	copy(result, nums)
	return result
}

func myappend(nums []int) []int {
	// 使用append函数
	var result []int
	result = append(result, nums...)
	return result
}

func main() {
	res := select3()
	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Println()

	sum := Sum(res)
	fmt.Printf("前三个分数的和为：%d", sum)
	fmt.Println()

	copyres := mycopy(res)
	copyres2 := myappend(res)
	for _, v := range copyres {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for _, v := range copyres2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
