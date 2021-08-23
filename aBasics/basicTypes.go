package main

type newType struct {
	field1 int
	field2 string
	field3 []bool
}

func funksjon(parameter1 int, parameter2 string, _ ubruktVerdi) returVerdi {

	retVal := funksjon(parameter1, parameter2, struct{}{})

	return retVal
}

func types() []interface{} {

	init, fun := medFlereReturns()

	intList := make([]int, 5) // liste, 5 lang


	initListe := []interface{}{init, fun, intList}
	return initListe
}

func medFlereReturns() (int, string) {
	return 2, "hei"
}
