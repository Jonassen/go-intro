package main

import (
	"fmt"
	"time"
)

// Keywords:
// break        default      func         interface    select
// case         defer        go           map          struct
// chan         else         goto         package      switch
// const        fallthrough  if           range        type
// continue     for          import       return       var

func controlFlow() {
	// While eq
	for {
		break
	}

	x := 0
	for x < 5 {
		x++
	}

	for i := 0; i < 10; i++ {
		x = i
	}

	if x == 0 {
		fmt.Println("X var 0")
	} else if false {
		fmt.Println("Never")
	} else {
		fmt.Println("X var ikke 0")
	}

	liste := make([]bool, 0)
	mapVar := make(map[int]bool)

	for index, value := range liste {
		fmt.Printf("Index: %v, value: %v\n", index, value)
	}

	for key, value := range mapVar {
		// Scrambled!
		fmt.Printf("Key: %v, value: %v\n", key, value)
	}

	s := "a"

	switch s {
	case "a":
		fmt.Println("a")
	case "b":
	case "c":
		fmt.Println("b or c")
	default:
		fmt.Println("Everything else")
	}

	secChan := time.After(1 * time.Second)
	fiveSecChan := time.After(5 * time.Second)

	select {
	case <-secChan:
		fmt.Println("1 seconds passed")
	case <-fiveSecChan:
		fmt.Println("5 seconds passed")
	}

	cleanUp := func() {
		fmt.Println("Clean up mess!")
	}

	defer cleanUp()

	fmt.Println("Create mess")
}
