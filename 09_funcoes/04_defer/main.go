// `defer` agenda uma chamada para o final da função atual e é muito usado para liberar recursos:

package main

import "fmt"

func primeira() {
	fmt.Println("Primeira")
}

func segunda() {
	fmt.Println("Segunda")
}

func ultima() {
	fmt.Println("Ultima")
}

func main() {
	defer ultima()
	primeira()
	segunda()
}
