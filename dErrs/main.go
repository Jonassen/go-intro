package main

import "fmt"

func main() {
 err := errorExample()
 if err != nil {
	 fmt.Printf("Error: %v\n", err.Error())
 }
}
