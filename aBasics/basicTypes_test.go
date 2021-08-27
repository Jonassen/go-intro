package main

import "testing"

func testbar() int {
	return 5
}

func TestTypes(t *testing.T) {
	if testbar() != 2 {
		t.Errorf("%v failed, testbar not 2", t.Name())
	}
	if testbar() != 3 {
		t.Fail()
	}
}
