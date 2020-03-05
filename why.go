package main

import "fmt"

var c chan string

func pingPang() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From PingPang: %d\n", i)
		i++
	}
}
func main() {
	c = make(chan string)
	go pingPang()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main: %d\n", i)
		fmt.Println(<-c)
	}
}
