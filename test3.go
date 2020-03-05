package main

import (
	"fmt"
	"strconv"
)

func toString() {
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("%T %v\n", str, str)
	fmt.Printf("Int to String")
}

func toInt() {
	str1 := "110"
	str2 := "s110"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("装换失败")
	} else {
		fmt.Println(num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("装换失败")
	} else {
		fmt.Println(num2)
	}

}

func toAppend() {
	b10 := []byte("")
	fmt.Println(b10)
	b10 = strconv.AppendInt(b10, 41, 10)
	fmt.Printf(string(b10))

}

func arr() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func newSlice() {

	var arr [30]int
	for i := 0; i < 30; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	fmt.Println(arr[1:5])
	fmt.Println(arr[5:7])
	fmt.Println(arr[28:len(arr)])
	fmt.Println(arr[:5])
}

func appendSlice() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Println(numbers)
	}
}
func test() {
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)
	fmt.Println(a)
	a = append([]int{-1, -2, -3}, a...)
	fmt.Println(a)
	var b = []int{5, 6}
	b = append(b, 8)
	fmt.Println(b)
	a = append(a[:len(a)], b...)
	fmt.Println(a)
}

func copyArray() {
	arr1 := []int{2, 3, 4, 5, 11, 99, 100}
	arr2 := []int{7, 8, 5, 80}
	fmt.Println(arr1[3])
	//copy(arr1, arr2)
	copy(arr1, arr2)
	fmt.Println(arr1)
}
func main() {
	copyArray()
	//test()
	//toString()
	//toInt()
	//toAppend()
	//arr()
	//newSlice()
	//appendSlice()
}
