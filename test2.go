package main

import "fmt"

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
	CORE
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	case CORE:
		return "CORE"
	}

	return "N/a"
}

func main(){

	fmt.Printf("%s %d",CPU,CORE)
}