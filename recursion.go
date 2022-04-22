package main

func Factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * Factorial(number-1)
	}

}
