package main

import (
	"fmt"
	"os"
)

func main() {
	tiposSorvete := map[string]float64{
		"casquinha": 1.00,
		"cascao":    2.50,
		"cestinha":  4.00,
	}

	sabores := map[string]float64{
		"morango":   0,
		"creme":     0,
		"chocolate": 0,
	}

	coberturas := map[string]float64{
		"caramelo":  1.50,
		"morango":   1.50,
		"chocolate": 1.50,
		"sem":       0.0,
	}

	itens := map[string]map[string]float64{
		"tipos":      tiposSorvete,
		"sabores":    sabores,
		"coberturas": coberturas,
	}

	var tipoSorvete, sabor, cobertura string

	fmt.Println("Tipo de Sorvete: casquinha (R$1,00), cascão (R$ 2,50), cestinha {R$ 4,00}")
	fmt.Scanf("%s", &tipoSorvete)

	fmt.Println("Sabor do Sorvete: morango, creme, chocolate")
	fmt.Scanf("%s", &sabor)

	fmt.Println("Cobertura: Caramelo (R$1,50), morango (R$ 1,50), chocolate {R$ 1,50} ou sem (R$ 0,00)")
	fmt.Scanf("%s", &cobertura)

	total := 0.00

	if valor, ok := itens["tipos"][tipoSorvete]; !ok {
		fmt.Println("Tipo de sorvete inválido")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := itens["sabores"][sabor]; !ok {
		fmt.Println("Sabor inválido")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := itens["coberturas"][cobertura]; !ok {
		fmt.Println("Cobertura inválida")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Printf("Valor a ser pago R$ %.2f", total)

}
