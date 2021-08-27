package main

import (
	"bytes"
	"fmt"
	"io"
)

type duck interface {
	quack(times int) string
	walk() string
}

type duck2 interface {
	quacker
	walker
}

type quacker interface {
	quack(times int) string
}

type walker interface {
	walk() string
}

type mallard struct{}

func (d mallard) quack(times int) string {
	retval := ""
	for i := 0; i < times; i++ {
		retval = retval + "Quack "
	}
	return ""
}

func (d mallard) walk() string {
	return "like a duck"
}

func doDuckStuff(d duck) {
	d.walk()
	d.quack(2)
}

func main() {
	var x io.ReadWriter = bytes.NewBuffer([]byte("data"))
	fmt.Println(x)

	m := mallard{}

	doDuckStuff(m)
}
