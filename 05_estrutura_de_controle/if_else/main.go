package main

import "fmt"

func main() {
	var idade int

	fmt.Println("Informe sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {
		fmt.Println("Você é idoso.")
	} else if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}
}
