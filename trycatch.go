package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func a() {
	fmt.Println("func a")
}
func b() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover b")
		}
	}()
	panic("panic b")
}
func c() {
	fmt.Println("func c")
}
