package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var notas [4]float64

	var nota string

	for i := 1; i <= len(notas); i++ {

		fmt.Printf("Informe a %d nota: ", i)
		fmt.Scanf("%s", &nota)
		nota, err := strconv.ParseFloat(nota, 64)

		if err != nil {
			fmt.Println("Erro ao converter nota")
		}
		notas[i-1] = nota
	}

	min := slices.Min(notas[:])
	max := slices.Max(notas[:])
	var total float64
	for _, v := range notas {
		total += v
	}

	fmt.Println("notas = ", notas)
	fmt.Println("Menor nota: ", min)
	fmt.Println("Maior nota: ", max)
	fmt.Println("Total de notas: ", total)
	fmt.Println("Média das notas: ", total/float64(len(notas)))
}
