package main

import "fmt"

type newType struct {
	privateField int
	PublicField  string
	PublicField2 []bool
}

func types() {
	integer := 5
	str := "String"
	integer = 2

	initialisert := newType{
		privateField: 2,
		PublicField:  "String",
		PublicField2: []bool{false, true},
	}

	initialisertTom := newType{}
	// Får tomme verdier:
	// int -> 0
	// bool -> false
	// string -> ""
	// lister -> []

	// Access
	initialisertTom.privateField = initialisert.privateField

	liste := []bool{false, false, true}
	liste[0] = true
	liste[1] = liste[2]

	mapVar := make(map[int]bool)
	mapVar[0] = false

	var pointer *int // nil
	pointer = &integer

	fmt.Println(str, integer, pointer)
}

func funksjon(parameter1 int, parameter2 string, _ ubruktVerdi) returVerdi {
	retVal := funksjon(parameter1, parameter2, struct{}{})

	init, fun := medFlereReturns()
	fmt.Println(init, fun)

	return retVal
}

func medFlereReturns() (int, string) {
	return 2, "hei"
}

type intList []int

type filter func(int) bool

func (i intList) filter(f filter) []int {
	retVal := make([]int, 0)
	for _, intFromList := range i {
		if f(intFromList) {
			retVal = append(retVal, intFromList)
		}
	}
	return retVal
}

func methodCall() {
	var iList intList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	greaterThan5 := func(i int) bool {
		return i > 5
	}

	_ = iList.filter(greaterThan5)
}
