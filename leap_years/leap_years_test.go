package main

import "testing"

func TestIsNotLeapWhenNotDivisibleBy4(t *testing.T) {
	year := Year(2001)
	result, _ := year.IsLeap()
	assertIsNotLeap(result, year, t)
}

func TestIsLeapWhenDivisibleBy4(t *testing.T) {
	year := Year(2004)
	result, _ := year.IsLeap()
	assertIsLeap(result, year, t)
}

func TestIsNotLeapWhenDivisibleBy100(t *testing.T) {
	year := Year(1900)
	result, _ := year.IsLeap()
	assertIsNotLeap(result, year, t)
}

func TestIsLeapWhenDivisibleBy400(t *testing.T) {
	year := Year(2000)
	result, _ := year.IsLeap()
	assertIsLeap(result, year, t)
}

func TestReturnsErrWhenBefore1583(t *testing.T) {
	year := Year(1500)
	_, err := year.IsLeap()
	if err == nil {
		t.Errorf("Rok %d jest sprzed okresu wprowadzenia kalendarza gregoriańskiego", year)
	}
}

func assertIsLeap(result bool, year Year, t *testing.T) {
	if result != true {
		t.Errorf("Rok %d jest przestępny, otrzymano: %v", year, result)
	}
}

func assertIsNotLeap(result bool, year Year, t *testing.T) {
	if result != false {
		t.Errorf("Rok %d jest przestępny, otrzymano: %v", year, result)
	}
}
