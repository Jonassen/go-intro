package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	f := false
	madMap := make(map[int]*bool)
	madMap[1] = &f

	fmt.Println(madMap[2])
}
