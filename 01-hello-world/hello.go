package helloworld

import "fmt"

func Hello() string {
	return "Hello, world!"
}

func main() {
	fmt.Print(Hello())
}
