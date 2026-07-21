// Exemplo fatia que recebe notas de um aluno

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var notas []float64
	notas := []float64{}

	for {
		var input string
		fmt.Println("Informe a nota do aluno")
		fmt.Scanf("%s", &input)

		if input == "" {
			break
		}
		nota, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Erro ao converter nota")
		}

		notas = append(notas, nota)

	}
	fmt.Printf("Notas: %v, tamanho len: %d\n", notas, len(notas))

}
