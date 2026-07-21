package main

import "fmt"

func main() {
	// Array vazio
	var x [10]float64
	fmt.Printf("x: %v\n", x)
	fmt.Println("-------------------")

	// Array iniciado
	t := [10]string{"Hello", "World"}
	fmt.Printf("t: %v\n", t)
	fmt.Printf("Primeiro elemento de t: %v\n", t[0])

	fmt.Println("-------------------")

	var y = [10]float64{0.32, 12, 42, 9.99}
	y[8] = 33
	y[9] = 44

	fmt.Printf("y: %v\n", y)
	fmt.Printf("Primeiro elemento de y: %v\n", y[0])
	fmt.Printf("Último elemento de y: %v\n", y[len(y)-1])
	fmt.Printf("Pegar os dois primeiros elementos de y: %v\n", y[:2])
	fmt.Printf("Tamanho do array y: %v\n", len(y))

	var total float64
	for i, v := range y {
		fmt.Println("Índice:", i, "Valor:", v)
		total += v
	}
	fmt.Println("Soma de todos os elementos de y:", total)

}
