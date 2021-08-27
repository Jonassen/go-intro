package main

import (
	"crypto/sha256"
	"fmt"
)

func errorExample() error {
	hash := sha256.New()

	var err error

	value, err := hash.Write([]byte("Data som hashes"))
	if err != nil {
		return err
	}
	// Now safe to use value

	bytes := hash.Sum(nil)
	fmt.Printf("%X", bytes)
	fmt.Println(value)

	return nil
}
