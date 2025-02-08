package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) float64 {
	if b == 0 {
		panic(errors.New("invalid value")) // 👈 panic declaration
	}
	return float64(a / b)
}

func main() {
	defer func() {
		r := recover() // 👈 recover
		if r != nil {
			fmt.Println("Panic occured", r)
			return
		}
		fmt.Println("Application running perfectly")
	}()

	// Test case 1
	a := 2.0
	b := 3.0
	fmt.Println(divide(a, b))

	// Test case 2
	b = 0.0
	fmt.Println(divide(a, b))
}
