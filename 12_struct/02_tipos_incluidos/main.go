package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint
	Peso   float64
	Altura float64
	Cpf    string
}

func (p Pessoa) GetIMC() float64 {
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

func (s *Smartphone) Ligar() {
	s.Ligado = true
}

func (s *Smartphone) Desligar() {
	s.Ligado = false
}

type Android struct {
	Smartphone
	OS string
}

// Método exclusivo do Android, não existe em Smartphone
func (an *Android) DownloadPlayStore() {
	fmt.Println("Baixando PlayStore")
}

type IOS struct {
	Smartphone
	OS string
}

// Método exclusivo do IOS, não existe em Smartphone e nem em Android
func (an *IOS) DownloadAppleStore() {
	fmt.Println("Baixando AppleStore")
}

func main() {
	s20 := Android{
		Smartphone: Smartphone{
			Marca:         "Samsung S20",
			Armazenamento: 256,
			Cor:           "Preto",
			Peso:          150.32,
			Valor:         3500.00,
		},
		OS: "Android 21",
	}

	fmt.Println(s20)
	s20.Ligar()
	fmt.Println("Ligando")
	fmt.Println("Ligado?", s20.Ligado)
	fmt.Println(s20)
	s20.DownloadPlayStore()

	fmt.Println("-----------")

	iphone := IOS{
		Smartphone: Smartphone{
			Marca:         "Apple",
			Armazenamento: 128,
			Cor:           "Prata",
			Peso:          130.23,
			Valor:         10000.00,
		},
		OS: "IOS X",
	}

	fmt.Println(iphone)
	iphone.Ligar()
	fmt.Println("Ligando")
	fmt.Println("Ligado?", iphone.Ligado)
	fmt.Println(iphone)
	iphone.DownloadAppleStore()

}
