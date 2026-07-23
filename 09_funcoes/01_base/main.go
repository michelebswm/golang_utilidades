package main

import (
	"errors"
	"fmt"
)

func soma(number1 int, number2 int) int {
	return number1 + number2
}

// Normal
// func media(number1, number2 int) float64 {
// 	return float64(soma(number1, number2)) / float64(2)
// }

// Retornando mais de um valor na função
func media(number1, number2 int) (float64, error) {
	return float64(soma(number1, number2)) / float64(2), fmt.Errorf("Erro")
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor não pode ser zero")
	}
	return a / b, nil
}

func main() {

	a, b := 10, 20

	resSoma := soma(a, b)
	fmt.Println("Soma: ", resSoma)

	resMedia , err := media(a, b)
	fmt.Println("Média: ", resMedia, err)

	c, d := 5.5, 0.0
	resDivisao, err := dividir(c, d)

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("Divisão: ", resDivisao)
	}

}
