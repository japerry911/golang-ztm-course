//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(name string) {
	fmt.Println("greetings", name)
}

func aMessage() string {
	return "this is a message"
}

func add3Ints(int1, int2, int3 int) int {
	return int1 + int2 + int3
}

func returnInt() int {
	return 32
}

func return2Ints() (int, int) {
	return 3, 2
}

func main() {
	greetPerson("Skylord")
	fmt.Println(aMessage())
	fmt.Println("1 + 2 + 3 =", add3Ints(1, 2, 3))
	fmt.Println("returnInt", returnInt())
	int1, int2 := return2Ints()
	fmt.Println("return2Ints =", int1, int2)
}
