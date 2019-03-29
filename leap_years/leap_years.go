package main

import "errors"

func IsLeap(year int) (bool, error) {
	if year < 1583 {
		return false, errors.New("Rok %d jest sprzed okresu wprowadzenia kalendarza gregoriańskiego")
	}
	switch {
	case year%400 == 0:
		return true, nil
	case year%100 == 0:
		return false, nil
	case year%4 == 0:
		return true, nil
	default:
		return false, nil
	}
}

func main() {
}
