package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// numbersToSum = []int{1, 2}, []int{1, 4}
	// resultado = [2]int{x, y}

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func main() {
	fmt.Println(Sum([]int{1, 2}))
	fmt.Println(Sum([]int{1, 4}))
	fmt.Println(SumAll([]int{1, 2}, []int{1, 4}))
}
