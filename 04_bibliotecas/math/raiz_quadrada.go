package main

import (
	"fmt"
	"math"
)

func main() {
	var num int

	fmt.Println("Informe um número inteiro: ")
	fmt.Scanf("%d", &num)

	raiz := math.Sqrt(float64(num))
	fmt.Printf("A raiz quadrada de %d é %.2f\n", num, raiz)

}
