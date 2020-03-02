package main

import "fmt"

func printArray(arr [5]int) {
	// [10]int 和 [20]int是不同类型
	// 调用func f(arr [10]int) 会拷贝数组
	// 值类型，一般不直接使用数组，而使用切片
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}          //冒号需要初值
	arr3 := [...]int{2, 4, 6, 8, 10} // 编译器决定长度
	var grid [4][5]int

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
