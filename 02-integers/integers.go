package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func ExampleAdder() {
	sum := Add(2, 2)
	fmt.Print(sum)
}
