package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}

func main() {
	fmt.Println(Sum([]int{1, 2}))
	fmt.Println(Sum([]int{1, 4}))
	fmt.Println(SumAllTails([]int{1, 2}, []int{1, 4}))
}
