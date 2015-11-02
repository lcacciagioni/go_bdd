package nextAfter

import "fmt"

// NextAfter is a function that returns the next number to 2 that you give
func NextAfter(number1, number2 int) (int, error) {
	if number1 > number2 {
		return 0, fmt.Errorf("number2 is bigger than number1")
	}
	return number1 + (2 * (number2 - number1)), nil
}
