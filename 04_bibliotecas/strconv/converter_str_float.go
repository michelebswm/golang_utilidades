package main

import (
	"fmt"
	"strconv"
)

func main() {
	var soma float64

	for {
		var entrada string
		fmt.Printf("Informe um número: ")
		fmt.Scanf("%s", &entrada)

		if entrada == ""{
			break
		}

		valor, err := strconv.ParseFloat(entrada, 64)

		if err != nil {
			fmt.Printf("Erro ao converter '%s' para float64: %v\n", entrada, err)
			continue
		}

		soma += valor
	}

	fmt.Printf("A soma dos números informados é: %f\n", soma)
}
