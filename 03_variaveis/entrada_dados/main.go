package main

import "fmt"

func main(){
	var x int
	fmt.Println("Informe um numero inteiro: ")
	fmt.Scanf("%d", &x) // Endereço na memória
	fmt.Println("O numero informado foi:", x)

	var y float64
	fmt.Println("Informe um numero real: ")
	fmt.Scanf("%f", &y)
	fmt.Println("O numero informado foi:", y)

	var z, w, v float64
	fmt.Println("Informe três números reais: ")
	fmt.Scanf("%f %f %f", &z, &w, &v)
	fmt.Println("Os números informados foram:", z, w, v)

	var nome string
	fmt.Println("Informe seu nome: ")
	fmt.Scanf("%s", &nome)
	fmt.Println("O nome informado foi:", nome)

	var ativo bool
	fmt.Println("Informe se está ativo (true/false): ")
	fmt.Scanf("%t", &ativo)
	fmt.Println("O status informado foi:", ativo)
}
