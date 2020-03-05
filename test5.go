package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make(map[string]string)
	arr["name"] = "jingzhen"
	arr["age"] = "31"
	arr["chineseName"] = "荆祯"

	var newArr []string

	for k := range arr {
		newArr = append(newArr, k)
	}
	fmt.Println(newArr)
	sort.Strings(newArr)
	fmt.Println(newArr)
	fmt.Println(arr)
	fmt.Println(arr["city"])
	if arr["city"] == "" {
		fmt.Println("无此健")
	}
	delete(arr, "city")
	delete(arr, "age")
	fmt.Println(arr)
	arr = make(map[string]string)
	fmt.Println(arr)

}
