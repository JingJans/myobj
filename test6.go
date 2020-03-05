package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("name", "荆祯")
	scene.Store("age", "31")
	scene.Store("address", "北京市海淀区")

	fmt.Println(scene.Load("address"))
	scene.Delete("address")
	i := 0
	scene.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		i++
		fmt.Println(i)
		return true
	})

}
