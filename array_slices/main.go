package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6} // este é um slice
	for _, number := range numbers {
		fmt.Println(number)
	}

	numbers = append(numbers, 10, 11, 12)
	fmt.Println(numbers)
}
