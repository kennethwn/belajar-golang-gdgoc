package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	/* Pointer explanation */
	var a int
	a = 10

	var b *int // 👈 this variable references to memory of variable a
	b = &a

	fmt.Println(b)
	fmt.Println(&a) // referencing technique
	fmt.Println(*b) // dereferencing technique

	a = 15

	fmt.Println(b)
	fmt.Println(*b)

	/* Simple case using pointer */
	num_1 := 5
	num_2 := 10
	swap(&num_1, &num_2)

	fmt.Println(num_1)
	fmt.Println(num_2)
}
