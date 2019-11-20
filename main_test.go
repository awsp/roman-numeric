package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	var notations = map[int]string{1000: "M", 500: "D", 100: "C", 50: "L", 10: "X", 5: "V", 1: "I",}

	actual := Calc(80, GetKeys(notations))
	expected := 4

	if len(actual) != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}

	fmt.Println(actual)
}

func TestNLength(t *testing.T) {
	actual := NLength(1202)
	expect := 4

	if actual != expect {
		t.Errorf("Expected: %d, Actual: %d", actual, expect)
	}
}

func TestConvert(t *testing.T) {
	var notations = map[int]string{1000: "M", 500: "D", 100: "C", 50: "L", 10: "X", 5: "V", 1: "I",}

	actual := Convert(Calc(98, GetKeys(notations)), notations)
	expect := "XCVIII"

	if actual != expect {
		t.Errorf("Expected: %s, Actual: %s", actual, expect)
	}
}