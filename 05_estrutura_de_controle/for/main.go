package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("--------------------")
	x := 0
	for x <= 5 {
		fmt.Printf("x = %d\n", x)
		x++
	}

	fmt.Println("--------------------")
	nome := "Michele BS"
	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
		// i => Indice
		// nome[i] => Byte
		// string(nome[i]) => String
	}

	fmt.Println("--------------------")
	nomes := []string{"Ana", "Bia", "Caio"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("--------------------")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("--------------------")
	// Break e continue

	name := "Michele BS"
	for _, v := range name {
		letra := string(v)

		fmt.Println(letra)

		if letra == "e" {
			fmt.Println("Encontrou a letra e!")
			continue
		}

		if letra == " " {
			break
		}
	}

}
