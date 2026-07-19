package main

import "fmt"

func main() {
	var x int
	fmt.Println("x =", x)

	var y int = 10
	fmt.Println("y =", y)

	x = 45
	fmt.Println("x =", x)

	var z, w, v float64
	fmt.Printf("z = %f | w = %f | v = %f \n", z, w, v)
	fmt.Printf("z = %T | w = %T | v = %T \n", z, w, v)

	
	var a, b, c string = "Estudo", "Golang", "today"
	fmt.Println(a, b, c)
}