package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	/* Closure or anonymous function */
	sum := 0
	summarize := func(numbers ...int) int {
		for _, num := range numbers {
			sum += num
		}
		return sum
	}

	fmt.Println(summarize(1, 2, 3, 4, 5))
	fmt.Println(summarize(1, 2, 3, 4, 5))

	/* Function as parameter (callback function) */
	printHelloWithName("Baju Kotor", customizeFilter)

	/* Function as value */
	sumarize := summary
	fmt.Println(sumarize(1, 2, 3))

	/* Variadic function */
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(summary(nums...))
	var employees = []string{"Joshua", "Kenneth", "Rayhan"}
	printList(employees, len(employees))

	/* Function return multiple value with named return value */
	luas, keliling := getAreaAndCircumference2(5)
	fmt.Printf("luas lingkaran dengan r adalah 5 = %.2f\n", luas)
	fmt.Printf("keliling lingkaran dengan r adalah 5 = %.2f\n", keliling)

	/* Function return multiple value */
	luas2, keliling2 := getAreaAndCircumference(5)
	fmt.Printf("luas lingkaran dengan r adalah 5 = %.2f\n", luas2)
	fmt.Printf("keliling lingkaran dengan r adalah 5 = %.2f\n", keliling2)

	/* Function return single value */
	fmt.Println(getFullName("Kenneth", "Noverianto"))

	/* Function (Parameter) */
	peoples := []string{"Joshua", "Kenneth", "Rayhan"}
	printList(peoples, len(peoples))

	/* function */
	generateHeader()
	generateBody()
}

func generateHeader() {
	fmt.Println("======================")
	fmt.Println("| Selamat Datang")
	fmt.Println("| di backend tutorial")
	fmt.Println("======================")
}

func generateBody() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d. Hello world\n", i)
	}
}

func printList(peoples []string, n int) {
	for _, people := range peoples {
		fmt.Println(people)
	}
	fmt.Println("Jumlah karyawan:", n)
}

func getFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func getAreaAndCircumference(r float64) (float64, float64) {
	area := math.Pi * r * r
	circumference := 2 * math.Pi * r
	return area, circumference
}

func getAreaAndCircumference2(r float64) (area, circumference float64) {
	area = math.Pi * r * r
	circumference = 2 * math.Pi * r
	return area, circumference
}

func summary(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func printHelloWithName(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func customizeFilter(text string) string {
	if strings.Contains(text, "kotor") {
		return "sensored"
	}
	return text
}
