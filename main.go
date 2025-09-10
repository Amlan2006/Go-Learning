package main

import "fmt"

func square(x int) int {
	return x * x
}
func sumAndProduct(x int, y int) (int, int) {
	return x + y, x * y
}
func main() {
	// fmt.Println("Hello")
	// var country = "India"
	// fmt.Println(" from", country)
	// language := "go"
	// fmt.Println("I love", language)
	// var x int
	// var y int
	// fmt.Println("Enter a number: ")
	// fmt.Scan(&x)
	// fmt.Println("You entered:", x)
	// fmt.Println("Enter a number: ")
	// fmt.Scan(&y)
	// fmt.Println("You entered:", y)
	// fmt.Println("Sum is:", x+y)
	// var array [5]string
	// fmt.Println(array)
	// loops("Hello")
	// conditions(9)
	var x int
	fmt.Println("Enter a number: ")
	fmt.Scan(&x)
	fmt.Println("Square is:", square(x))
	a, b := sumAndProduct(x, 5)
	fmt.Println("Sum and Product with 5 is:", a, b)

}
