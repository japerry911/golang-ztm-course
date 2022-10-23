package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet func")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("dozen", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("bakersDozen", bakersDozen)
}
