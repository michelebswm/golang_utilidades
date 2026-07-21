package main

import "fmt"

func main() {
	// Criando uma fatia
	s := []int{1, 2, 3}
	fmt.Println("Slice:", s)

	// Adicionando elementos
	s = append(s, 4, 5)
	fmt.Println("Slice após adicionar elementos:", s)

	// Removendo elementos
	s = s[:len(s)-1]
	fmt.Println("Slice após remover o último elemento:", s)

	// Iterando sobre a fatia
	for i, v := range s {
		fmt.Println("Índice:", i, "Valor:", v)
	}

	fmt.Println("-------------------")
	fmt.Println("Fatia pode aumentar ou diminuir de tamanho, já o Array não.")
	var x [100]int
	fatia := x[:10]
	fmt.Printf("Fatia: %v\n", fatia)
	fmt.Printf("X tipo %T Fatia tipo: %T, tamanho len: %d\n", x, fatia, len(fatia))

	fmt.Println("-------------------")
	// Criando uma nova fatia
	y := []int{}
	fmt.Printf("Y tipo: %T, tamanho len: %d\n", y, len(y))

	fmt.Println("-------------------")
	for i := 1; i<=200;i++{
		y = append(y, i)
		fmt.Println(len(y))
	}
	fmt.Printf("Y: %v\n, capacidade: %d\n", y, cap(y))
}
