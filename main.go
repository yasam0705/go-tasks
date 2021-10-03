package main

import (
	"fmt"
	"udevs-tasks/tasks"
)

func main() {
	// tasks.FizzBuzz(16)

	// fmt.Println(tasks.Fibonacci(8))

	// tasks.Palindrome()

	// fmt.Println(tasks.OddEven([]int{1, 3, 8, 4, 6}))

	fmt.Println(tasks.HasDuplicate([]int{1, 3, 8, 4, 6, 1}))
}
