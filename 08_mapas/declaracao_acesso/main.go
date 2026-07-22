package main

import "fmt"

func main() {
	// Criando um map com make
	idades := make(map[string]int)

	idades["Michele"] = 35
	idades["Wallace"] = 42

	fmt.Println("Idades", idades)
	fmt.Println("Tamanho do map:", len(idades))

	// Outra forma de iniciar map
	alturas := map[string]float64{}

	alturas["Michele"] = 1.62

	fmt.Println("Alturas", alturas)

	//--------------

	// Percorrendo um map
	for k, v := range idades {
		fmt.Println(k, v)
	}

	// removendo chaves
	delete(idades, "Michele")

	// Acessando um valor que existe
	alturaMi := alturas["Michele"]
	fmt.Println("Alturaa Michele: ", alturaMi)

	// Acessando um valor que não existe
	alturaNil, ok := alturas["teste"]
	fmt.Println(ok)
	if ok {
		fmt.Println("Altura alturaNil: ", alturaNil)
	} else {
		fmt.Println("Valor não existe")
	}

	//alturas["lara"] = 1.20
	// Variável valor existe somente dentro do if 
	if valor, ok := alturas["lara"]; ok{
		fmt.Println("A altura da lara é ", valor)
	} else {
		fmt.Println("Valor não existe")
	}

}
