package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func chengfa() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func switchCase(a string) int {
	switch a {
	case "":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	default:
		return 5
	}
}

func testGoto() {
	for x := 0; x <= 5; x++ {
		for y := 0; y <= 5; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("do other things!!!")
}

func robot() {

	inputHeader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字：")
	input, err := inputHeader.ReadString('\n')
	if err != nil {
		fmt.Printf("错误：%s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-2]
		fmt.Printf("你好！ %s! 我能为你做点什么？", name)
	}
	for {
		input, err := inputHeader.ReadString('\n')
		if err != nil {
			fmt.Printf("异常错误：%s\n", err)
			continue
		}
		input = strings.ToLower(input[:len(input)-2])
		fmt.Println(input)
		switch input {
		case "":
			continue
		case "bye", "fuck":
			fmt.Println("很高兴能帮到你,再见")
			os.Exit(0)
		default:
			fmt.Println("我不能明白您的意思？")
		}
	}

}

func main() {
	robot()
	//testGoto()
	//chengfa()
	//result := switchCase("aa")
	//fmt.Println(result)

}
