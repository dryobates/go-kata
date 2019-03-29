package main

import "testing"

func TestIsNotLeapWhenNotDivisibleBy4(t *testing.T) {
	year := 2001
	result, _ := IsLeap(year)
	assertIsNotLeap(result, year, t)
}

func TestIsLeapWhenDivisibleBy4(t *testing.T) {
	year := 2004
	result, _ := IsLeap(year)
	assertIsLeap(result, year, t)
}

func TestIsNotLeapWhenDivisibleBy100(t *testing.T) {
	year := 1900
	result, _ := IsLeap(year)
	assertIsNotLeap(result, year, t)
}

func TestIsLeapWhenDivisibleBy400(t *testing.T) {
	year := 2000
	result, _ := IsLeap(year)
	assertIsLeap(result, year, t)
}

func TestReturnsErrWhenBefore1583(t *testing.T) {
	year := 1500
	_, err := IsLeap(year)
	if err == nil {
		t.Errorf("Rok %d jest sprzed okresu wprowadzenia kalendarza gregoriańskiego", year)
	}
}

func assertIsLeap(result bool, year int, t *testing.T) {
	if result != true {
		t.Errorf("Rok %d jest przestępny, otrzymano: %v", year, result)
	}
}

func assertIsNotLeap(result bool, year int, t *testing.T) {
	if result != false {
		t.Errorf("Rok %d jest przestępny, otrzymano: %v", year, result)
	}
}
