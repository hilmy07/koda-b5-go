package main

import (
	"errors"
	"fmt"
)

func ProcessNumber(numbers []int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovery from Panic: %v\n", r)
		}
	}()

	// fmt.Println(numbers == nil)

	if numbers == nil {
		return errors.New("No data provided")
	}

	if len(numbers) == 0 {
		panic("empty list provided")
	}

	for _, n := range numbers {
		fmt.Println(n * 2)
	}

	return nil
}

func main() {
	ProcessNumber([]int{3, 2})
	ProcessNumber([]int{})
	fmt.Println(ProcessNumber(nil))
}

