package main

import "fmt"

func soma(values ...int) int {

	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	a, b, c, d := 10, 20, 30, 40

	total := soma(a, b, c, d)
	fmt.Println(total)

	valor := []int{1,50,2,90,50,66,58,8,10,16,}
	result := soma(valor...)
	fmt.Println("Soma", result)

}
