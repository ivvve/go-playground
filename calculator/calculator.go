package calculator

import "errors"

func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Divide(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return float64(num1) / float64(num2), nil
}
