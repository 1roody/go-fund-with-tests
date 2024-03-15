package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	sum := Add(2, 2)
	fmt.Print(sum)
}
