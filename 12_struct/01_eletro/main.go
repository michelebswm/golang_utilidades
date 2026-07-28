package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade uint
	Peso  float64
	Altura  float64
	Cpf string
}

func (p Pessoa) GetIMC() float64{
	return p.Peso / (p.Altura * p.Altura)
}

type Smartphone struct {
	Marca         string
	Armazenamento uint
	Cor           string
	Peso          float64
	Valor         float64
	Ligado        bool
	Proprietario  Pessoa
}

// Método que altera o próprio objeto
// Métodos com letra Maiusculas são publicos, já com letra minuscula são privados no proprio package
func (s *Smartphone) Ligar() {
	s.Ligado = true
}

func (s *Smartphone) Desligar() {
	s.Ligado = false
}

func main() {

	mi := Pessoa{
		Nome: "Michele",
		Idade: 35,
		Peso: 85.5,
		Altura: 1.62,
		Cpf: "00000000",
	}

	celMi := Smartphone{
		Marca:         "Samsung S20",
		Armazenamento: 256,
		Cor:           "Preto",
		Peso:          150.32,
		Valor:         3500.00,
		Proprietario: mi,
	}

	fmt.Println(celMi)

	fmt.Println("Marca: ", celMi.Marca)
	fmt.Println("Armazenamento: ", celMi.Armazenamento)
	fmt.Println("Cor: ", celMi.Cor)
	fmt.Println("Peso: ", celMi.Peso)
	fmt.Println("Valor: ", celMi.Valor)
	fmt.Println("Ligado?: ", celMi.Ligado)
	fmt.Println("Nome do Proprietario: ", celMi.Proprietario.Nome)
	fmt.Println("Idade do Proprietario: ", celMi.Proprietario.Idade)


	fmt.Println("Ligando smartphone")
	celMi.Ligar()
	fmt.Println("Ligado?: ", celMi.Ligado)

	fmt.Println("Desligando smartphone")
	celMi.Desligar()
	fmt.Println("Ligado?: ", celMi.Ligado)

	// Alterando valores
	celMi.Peso += 20
	fmt.Println("Novo Peso: ", celMi.Peso)
	fmt.Println(celMi)

	fmt.Println("IMC do Proprietário: ", celMi.Proprietario.GetIMC())

}
