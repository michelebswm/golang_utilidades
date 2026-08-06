package main

import "fmt"

func dobrar(entrada <-chan int, saida chan<- int) {
	for numero := range entrada {
		saida <- numero * 2
	}
	close(saida)
}

func main() {
	entrada := make(chan int)
	saida := make(chan int)

	go dobrar(entrada, saida)

	go func() {
		defer close(entrada)
		for _, numero := range []int{1, 2, 3} {
			entrada <- numero
		}
	}()

	for resultado := range saida {
		fmt.Println(resultado)
	}
}

/*
O fluxo desse programa é:

1. `main` cria os canais `entrada` e `saida`.
2. A primeira goroutine recebe números de `entrada`, dobra cada um e envia para `saida`.
3. A função anônima envia `1`, `2` e `3`, e fecha `entrada` ao terminar.
4. O fechamento faz o `range entrada` terminar.
5. `dobrar` fecha `saida`, fazendo o `range saida` em `main` terminar.
*/