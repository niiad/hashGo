package base

import "fmt"

// AreaOfCircle calculates the area of a circle given its radius.
//
// Parameters:
//
//	radius: The radius of the circle.
//
// Returns:
//
//	The area of the circle as a float64.
//	An error if the input radius is negative.
func AreaOfCircle(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("radius must be non-negative")
	}

	area := 3.14159 * radius * radius

	return area, nil
}

// IterativeFactorial returns the factorial of a non-negative number using iteration
//
// Parameters:
//
//	number: the non-negative number for which to find the factorial
//
// Returns:
//
//	the factorial of the input number as an int
func IterativeFactorial(number int) int {
	if number < 0 {
		panic("number must be non-negative")
	}

	result := 1

	for i := 2; i <= number; i++ {
		result *= i
	}

	return result
}

// RecursiveFactorial returns the factorial of a non-negative number using recursion
//
// Parameters:
//
//	number: the non-negative number for which to find the factorial
//
// Returns:
//
//	the factorial of the input number as an int
func RecursiveFactorial(number int) int {
	if number < 0 {
		panic("number must be non-negative")
	}

	if number == 0 || number == 1 {
		return 1
	}

	return number * RecursiveFactorial(number-1)
}

// TailRecursiveFactorial returns the factorial of a non-negative number using tail recursion
//
// Parameters:
//
//	number: the non-negative number for which to find the factorial
//
// Returns:
//
//	the factorial of the input number as an int
func TailRecursiveFactorial(number int, accumulator int) int {
	if number < 0 {
		panic("number must be non-negative")
	}

	if number == 0 {
		return accumulator
	}

	return TailRecursiveFactorial(number-1, number*accumulator)

}
