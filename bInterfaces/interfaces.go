package main

import "fmt"

type and interface{
	kvekk(ganger int)
}

type donald struct{
	antallBen int
}

func (d donald) kvekk() {
	panic("implement me")
}

func gjørAndeTing(andemor and) {
	andemor.kvekk(2)
}

func main() {
	x := 4

	fmt.Println("hekkan")
}
