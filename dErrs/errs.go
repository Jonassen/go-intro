package main

import (
	"crypto/sha256"
	"fmt"
)

func errorExample() error {
	hash := sha256.New()

	_, err := hash.Write([]byte("Data som hashes"))
	if err != nil {
		return err
	}

	bytes := hash.Sum(nil)
	fmt.Printf("%X", bytes)

	return nil
}
